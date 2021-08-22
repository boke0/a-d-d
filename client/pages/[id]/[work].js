import styles from '../../styles/Work.module.scss';
import Header from '../../components/org/Header';
import axios from '../../axios';

export default function Work({ isLoggedIn, loginUser, work }) {
  const duration = new Date(Date.parse(work.EndTime) - Date.parse(work.StartTime));
  return (
    <div className={styles.container}>
      <Header isLoggedIn={isLoggedIn} loginUser={loginUser} />
      <main>
        <div className={styles.linedProfile}>
          <div className={styles.icon}>
            <img src={work.User.Icon} />
          </div>
          {work.User.Name}
        </div>
        <h4>{work.Title}</h4>
        <div className={styles.details}>
          <div className={styles.status}>
            <dl>
              <dt>作業時間</dt>
              <dd className={styles.currentTime}>{`${("00" + duration.getUTCHours()).slice(-2)}:${("00" + duration.getUTCMinutes()).slice(-2)}:${("00" + duration.getUTCSeconds()).slice(-2)}`}</dd>
              <dt>現在のアルコール摂取量</dt>
              <dd>{work.Drinks.reduce((value, drink) => (value + 0.0008 * drink.Alcohol * drink.Amount), 0)}g ({work.Drinks.length}杯目)</dd>
            </dl>
          </div>
        </div>
        <div>
          <h2>進捗</h2>
          <div>+{work.Add} -{work.Remove} ({work.Commits}コミット)</div>
        </div>
        <div className={styles.drinks}>
          <h2>飲んだお酒</h2>
          <div className={styles.drinkList}>
            {work.Drinks.map(({ Name, Alcohol, Amount }, i) => (
              <div className={styles.drinkItem}>
                <span>{i+1}杯目</span>
                <h4>{Name}</h4>
                <p>度数: {Alcohol}%, 量: {Amount}ml</p>
              </div>
            ))}
          </div>
        </div>
      </main>
    </div>
  )
}

export async function getServerSideProps({ params }) {
  const { work } = params;
  const work_ = await axios.get('/works/' + work).then(res => {
    return res.data.Work
  });
  return {
    props: {
      work: work_
    }
  }
}
