import { useEffect } from 'react';
import { useRouter } from 'next/router';
import { setCookie, destroyCookie } from 'nookies';
import axios from '../axios';

export default function Auth(ctx) {
  const router = useRouter()
  useEffect(() => {
    if(process.browser){
      const qs = new URL(location.href).searchParams;
      axios.post('/login', { code: qs.get("code") }).then(res => {
        setCookie(ctx, 'access_token', res.data.token, {
          maxAge: 3600 * 24 * 14
        })
        location.href = '/'
      })
    }
  }, []);
  return (
    <div>
      GitHubと連携中...
    </div>
  )
}
