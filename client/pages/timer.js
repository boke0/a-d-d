import { useState, useEffect, useRef } from 'react';
import { useRouter } from 'next/router';
import Btn from '../components/atm/Btn'
import styles from '../styles/Timer.module.scss'
import axios from '../axios';

const PREPARE = 0;
const IN_PROGRESS = 1;

export default function Timer({ isLoggedIn, loginUser }) {
  const router = useRouter();
  const [phase, setPhase] = useState(0);
  const startTime = useRef(new Date());
  const workId = useRef(null);
  const [currentTime, setCurrentTime] = useState('');
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [drinkName, setDrinkName] = useState('');
  const [drinkAlcohol, setDrinkAlcohol] = useState('');
  const [drinkAmount, setDrinkAmount] = useState('');
  const [drinks, setDrinks] = useState([]);
  const [open, setOpen] = useState(false);
  const [checking, setChecking] = useState(false);
  const updater = useRef(null);
  const start = () => {
    axios.post('/works', {
      title: title,
      description: description,
      start_time: new Date().toISOString(),
      drinks: [
        {
          name: drinkName,
          alcohol: Number(drinkAlcohol),
          amount: Number(drinkAmount)
        }
      ]
    }).then(res => {
      setPhase(IN_PROGRESS);
      startTime.current = new Date();
      setDrinks(res.data.Work.Drinks)
      setTitle(res.data.Work.Title)
      setDescription(res.data.Work.Description)
      workId.current = res.data.Work.ID;
    });
  }
  const postdrink = () => {
    axios.post('/works/' + workId.current + '/drinks', {
      name: drinkName,
      alcohol: Number(drinkAlcohol),
      amount: Number(drinkAmount)
    }).then(res => {
      setDrinks(drinks => ([
        ...drinks,
        res.data.Drink
      ]))
    });
    setOpen(false);
  }
  const openmodal = () => {
    setOpen(true);
  }
  const close = () => {
    setDrinkName(null)
    setDrinkAlcohol(null)
    setDrinkAmount(null)
    setOpen(false);
  }
  const end = () => {
    axios.put('/works/' + workId.current, {
      end_time: new Date().toISOString(),
      title: title,
      description: description,
    }).then(res => {
      router.push('/works/' + workId.current)
    });
  }
  useEffect(() => {
    if(!isLoggedIn){
      location.href = `https://github.com/login/oauth/authorize?scope=user,repo&client_id=${process.env.GITHUB_CLIENT_ID}`;
    }
    axios.get('/works/in_progress').then(res => {
      setPhase(IN_PROGRESS);
      const [ymd, his] = res.data.Work.StartTime.split('T');
      const [y, m, d] = ymd.split('-'), [h, i, s] = his.split('Z')[0].split(':')
      const date = new Date();
      date.setUTCFullYear(Number(y))
      date.setUTCMonth(Number(m) - 1)
      date.setUTCDate(Number(d))
      date.setUTCHours(Number(h))
      date.setUTCMinutes(Number(i))
      date.setUTCSeconds(Number(s))
      startTime.current = date;
      setDrinks(res.data.Work.Drinks)
      setTitle(res.data.Work.Title)
      setDescription(res.data.Work.Description)
      workId.current = res.data.Work.ID;
    }).catch(() => {
      setPhase(PREPARE);
    });
    updater.current = setInterval(() => {
      const diff = new Date(new Date().getTime() - startTime.current.getTime())
      setCurrentTime(`${diff.getUTCHours()}:${("00"+diff.getUTCMinutes()).slice(-2)}:${("00"+diff.getUTCSeconds()).slice(-2)}`)
    }, 1000)
  }, []);
  switch(phase){
    case PREPARE:
      return (
        <div className={styles.container}>
          <div className={styles.drinkform}>
            <h2 className={styles.title}>???????????????????????????????????????</h2>
            <input type='text' value={title} onInput={(e) => setTitle(e.target.value)} placeholder='?????????????????????' />
            <textarea onInput={(e) => setDescription(e.target.value)} placeholder='????????????' value={description}></textarea>
          </div>
          <div className={styles.drinkform}>
            <h2 className={styles.title}>????????????????????????????????????</h2>
            <input type='text' value={drinkName} onInput={(e) => setDrinkName(e.target.value)} placeholder='???????????????' />
            <input type='number' value={drinkAlcohol} onInput={(e) => setDrinkAlcohol(e.target.value)} placeholder='?????????????????????(%)'/>
            <input type='number' value={drinkAmount} onInput={(e) => setDrinkAmount(e.target.value)} placeholder='???(ml)'/>
            <div className={styles.submit}>
              <Btn onClick={start}>??????????????????</Btn>
            </div>
          </div>
        </div>
      )
    case IN_PROGRESS:
      return (
        <>
          <div className={styles.container}>
            <h1>????????????????????????</h1>
            <div className={styles.timer}>
              <div className={styles.currentTime}>{currentTime}</div>
              <div className={styles.status}>
                <dl>
                  <dt>?????????????????????????????????</dt>
                  <dd>{drinks.reduce((value, drink) => (value + 0.0008 * drink.Alcohol * drink.Amount), 0)}g ({drinks.length}??????)</dd>
                </dl>
              </div>
              <div className={styles.buttons}>
                <Btn onClick={openmodal} style={{ marginRight: 8 }}>???????????????</Btn>
                {/*<Btn onClick={e => setChecking(true)}>????????????</Btn>*/}
                <Btn onClick={end}>????????????</Btn>
              </div>
            </div>
            <div className={styles.drinks}>
              <h2>???????????????</h2>
              <div className={styles.drinkList}>
                {drinks.map(({ Name, Alcohol, Amount }, i) => (
                  <div className={styles.drinkItem}>
                    <span>{i+1}??????</span>
                    <h4>{Name}</h4>
                    <p>??????: {Alcohol}%, ???: {Amount}ml</p>
                  </div>
                ))}
              </div>
            </div>
          </div>
          <div className={`${styles.modal} ${open ? styles.show : ''}`}>
            <div className={styles.overlay} onClick={close}></div>
            <div className={`${styles.drinkform} ${styles.modalInner}`}>
              <h2 className={styles.title}>{title}</h2>
              <input type='text' value={drinkName} onInput={(e) => setDrinkName(e.target.value)} placeholder='???????????????' />
              <input type='number' value={drinkAlcohol} onInput={(e) => setDrinkAlcohol(e.target.value)} placeholder='?????????????????????(%)'/>
              <input type='number' value={drinkAmount} onInput={(e) => setDrinkAmount(e.target.value)} placeholder='???(ml)'/>
              <div className={styles.submit}>
                <Btn onClick={close} style={{ marginRight: 8 }} outline>???????????????</Btn>
                <Btn onClick={postdrink}>??????????????????</Btn>
              </div>
            </div>
          </div>
          {/*
          <div className={`${styles.modal} ${checking ? styles.show : ''}`}>
            <div className={styles.overlay} onClick={close}></div>
            <div className={`${styles.drinkform} ${styles.modalInner}`}>
              <h3>Git?????????????????????????????????</h3>
              <p>??????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????</p>
              <div className={styles.submit}>
                <Btn onClick={e => setChecking(false)} style={{ marginRight: 8 }} outline>??????????????????</Btn>
                <Btn onClick={end}>????????????????????????</Btn>
              </div>
            </div>
          </div>
          */}
        </>
      )
  }
}
