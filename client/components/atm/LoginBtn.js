import Btn from './Btn';
import { Github } from '@material-ui/icons';

export default function LoginBtn() {
  return (
    <a href={`https://github.com/login/oauth/authorize?scope=user,repo&client_id=${process.env.GITHUB_CLIENT_ID}`}>
      <Btn style={{
        backgroundColor: '#171515'
      }}>GitHubでログイン</Btn>
    </a>
  )
}
