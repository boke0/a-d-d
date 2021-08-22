import styles from '../../styles/org/Header.module.scss';
import Link from 'next/link'
import Btn from '../atm/Btn'
import LoginBtn from '../atm/LoginBtn'

export default function Header({ isLoggedIn, loginUser, children }) {
  return (
    <div className={styles.container}>
      <Link href='/'><h1>a-d-d.life</h1></Link>
      <div className={styles.left}>
      {
        isLoggedIn
          ? (
            <Link href='/timer'>
              <Btn>タイマー</Btn>
            </Link>
          )
          : (
            <LoginBtn />
          )
      }
      {
        loginUser != null
          ? (
            <Link href={`/${loginUser.ScreenName}`}>
              <div className={styles.profile}>
                <img src={loginUser.Icon} className={styles.icon} />
                <div>
                  {loginUser.Name}
                </div>
              </div>
            </Link>
          )
          : (<></>)
      }
      </div>
    </div>
  )
}
