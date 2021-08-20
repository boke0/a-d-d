import Head from 'next/head'
import Link from 'next/link'
import Image from 'next/image'
import styles from '../styles/Home.module.scss'
import Btn from '../components/atm/Btn'
import LoginBtn from '../components/atm/LoginBtn'

export default function Home({ isLoggedIn }) {
  return (
    <div className={styles.container}>
      <div className={styles.headers}>
        <h1>a-d-d.life</h1>
        <div>
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
        </div>
      </div>
      <main>
        <h1>みんなのADD</h1>
      </main>
    </div>
  )
}
