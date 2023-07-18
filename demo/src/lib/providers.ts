import { GithubAuthProvider, GoogleAuthProvider } from 'firebase/auth'

const googleAuthProvider = new GoogleAuthProvider()
googleAuthProvider.addScope('email')
googleAuthProvider.addScope('profile')

const githubAuthProvider = new GithubAuthProvider()

export { githubAuthProvider, googleAuthProvider }
export default { googleAuthProvider, githubAuthProvider }
