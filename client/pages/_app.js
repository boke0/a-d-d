import '../styles/globals.scss'
import { parseCookies } from 'nookies';
import axios from '../axios';

function MyApp(ctx) {
  const cookie = parseCookies(ctx);
  if(cookie.access_token) axios.defaults.headers.common['Authorization'] = cookie.access_token
  const isLoggedIn = cookie.access_token != null
  const { Component, pageProps } = ctx;
  return <Component isLoggedIn={isLoggedIn} {...pageProps } />
}

export default MyApp
