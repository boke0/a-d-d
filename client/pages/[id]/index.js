import styles from '../../styles/User.module.scss';
import Header from '../../components/org/Header';
import Link from 'next/link';
import axios from '../../axios';

export default function User({ isLoggedIn, loginUser, user }) {
  return (
    <div className={styles.container}>
      <Header isLoggedIn={isLoggedIn} loginUser={loginUser} />
      <main>
        <div className={styles.profile}>
          <div className={styles.icon}>
            <img src={user.Icon} />
          </div>
          <div className={styles.detail}>
            <h1 className={styles.name}>{ user.Name }</h1>
            <div className={styles.screenName}>{ user.ScreenName }</div>
            <p className={styles.description}>{ user.Description }</p>
          </div>
        </div>
        <div className={styles.works}>
          <h2>過去のADD</h2>
          <div className={styles.workList}>
            {user.Works.map((work, i) => {
              const duration = new Date(Date.parse(work.EndTime) - Date.parse(work.StartTime));
              return (
                <Link href={`/${user.ScreenName}/${work.ID}`} key={i}>
                    <div className={styles.workItem}>
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
  const { id } = params;
  const user = await axios.get('/users/' + id).then(res => {
    return res.data.User
  });
  return {
    props: {
      user
    }
  }
}
