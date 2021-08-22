import { useState, useEffect } from 'react';
import '../styles/globals.scss'
import { parseCookies } from 'nookies';
import axios from '../axios';

function MyApp(ctx) {
  const [loginUser, setLoginUser] = useState(null);
  const cookie = parseCookies(ctx);
  if(cookie.access_token) axios.defaults.headers.common['Authorization'] = cookie.access_token
  const isLoggedIn = cookie.access_token != null
  useEffect(() => {
    if(isLoggedIn) {
      axios.get('/session').then(res => {
        setLoginUser(res.data.User)
      })
    }
  }, [isLoggedIn])
  const { Component, pageProps } = ctx;
  return <Component isLoggedIn={isLoggedIn} loginUser={loginUser} {...pageProps } />
}

export default MyApp
