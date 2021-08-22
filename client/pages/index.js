import Head from 'next/head'
import Link from 'next/link'
import Image from 'next/image'
import styles from '../styles/Home.module.scss'
import Btn from '../components/atm/Btn'
import LoginBtn from '../components/atm/LoginBtn'
import Header from '../components/org/Header';
import axios from '../axios';

export default function Home({ isLoggedIn, loginUser, works }) {
  return (
    <div className={styles.container}>
      <Header isLoggedIn={isLoggedIn} loginUser={loginUser} />
      <main>
        <div className={styles.works}>
          <h1>みんなのADD</h1>
          <div className={styles.workList}>
            {works.map((work, i) => {
              const duration = new Date(Date.parse(work.EndTime) - Date.parse(work.StartTime));
              return (
                <Link href={`/${work.User.ScreenName}/${work.ID}`} key={i}>
                    <div className={styles.workItem}>
                      <div className={styles.profile}>
                        <img src={work.User.Icon} />
                        <div className={styles.name}>{work.User.Name}</div>
                      </div>
                      <h4>{work.Title}</h4>
                      <div>作業時間: {`${("00" + duration.getUTCHours()).slice(-2)}:${("00" + duration.getUTCMinutes()).slice(-2)}:${("00" + duration.getUTCSeconds()).slice(-2)}`}</div>
                      <div>{work.Drinks.length}杯</div>
                      <div>総アルコール量: {work.Drinks.reduce((value, drink) => (value + 0.0008 * drink.Alcohol * drink.Amount), 0)}g</div>
                    </div>
                </Link>
              )
            })}
          </div>
        </div>
      </main>
    </div>
  )
}

export async function getServerSideProps({ params }) {
  const works = await axios.get('/works').then(res => {
    return res.data.Works
  });
  return {
    props: {
      works
    }
  }
}
