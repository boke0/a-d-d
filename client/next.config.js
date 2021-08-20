const path = require('path')

let API_BASE_URL;
if(process.env.NODE_ENV == 'development'){
  API_BASE_URL = 'http://localhost:8000/v1'
}else{
  API_BASE_URL = 'https://api.a-d-d.life'
}

module.exports = {
  env: {
    GITHUB_CLIENT_ID: process.env.GITHUB_CLIENT_ID,
    API_BASE_URL,
  },
  sassOptions: {
    includePaths: [path.join(__dirname, 'styles')]
  }
}
