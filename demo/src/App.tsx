import {
  createUserWithEmailAndPassword,
  onAuthStateChanged,
  signInWithEmailAndPassword,
  signInWithPopup,
  signOut,
  User,
} from 'firebase/auth'
import { FormEvent, useEffect, useReducer, useState } from 'react'
import './App.css'
import { auth } from './lib/firebase'
import { githubAuthProvider, googleAuthProvider } from './lib/providers'

type Form = {
  email: string
  password: string
}
type FormReducer = (prev: Form, next: Form) => Form

const SignInForm = () => {
  const [loading, setLoading] = useState(false)

  const [form, updateForm] = useReducer<FormReducer>(
    (prev, next) => {
      return { ...prev, ...next }
    },
    { email: '', password: '' }
  )

  const onSignUp = async (e: FormEvent) => {
    e.preventDefault()
    setLoading(true)

    const credentials = await signInWithEmailAndPassword(
      auth,
      form.email,
      form.password
    )
      .catch((error) => {
        console.log(error)
      })
      .finally(() => setLoading(false))

    console.log(credentials)
  }

  return (
    <form onSubmit={(e) => onSignUp(e)}>
      <div className="input-group">
        <label htmlFor="email">E-Mail Address</label>
        <input
          id="email"
          type="email"
          placeholder="john@gmail.com"
          value={form.email}
          onChange={(e) => updateForm({ ...form, email: e.target.value })}
        />
      </div>
      <div className="input-group">
        <label htmlFor="password">Password</label>
        <input
          id="password"
          type="password"
          placeholder="password"
          value={form.password}
          onChange={(e) => updateForm({ ...form, password: e.target.value })}
        />
      </div>
      <button type="submit" className="submit-btn" disabled={loading}>
        Sign In
      </button>
    </form>
  )
}

const SignUpForm = () => {
  const [loading, setLoading] = useState(false)

  const [form, updateForm] = useReducer<FormReducer>(
    (prev, next) => {
      return { ...prev, ...next }
    },
    { email: '', password: '' }
  )
  const onSignUp = async (e: FormEvent) => {
    e.preventDefault()
    setLoading(true)

    const credentials = await createUserWithEmailAndPassword(
      auth,
      form.email,
      form.password
    )
      .catch((error) => {
        console.log(error)
      })
      .finally(() => setLoading(false))

    // send user data to custom backend api
    const user = {
      uid: credentials?.user?.uid,
      email: credentials?.user?.email,
    }

    const token = await credentials?.user?.getIdToken()

    const response = await fetch('http://localhost:9000/sign-in', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(user),
    })

    console.log(response)
  }

  return (
    <form onSubmit={(e) => onSignUp(e)}>
      <div className="input-group">
        <label htmlFor="email">E-Mail Address</label>
        <input
          id="email"
          type="email"
          placeholder="john@gmail.com"
          value={form.email}
          onChange={(e) => updateForm({ ...form, email: e.target.value })}
        />
      </div>
      <div className="input-group">
        <label htmlFor="password">Password</label>
        <input
          id="password"
          type="password"
          placeholder="password"
          value={form.password}
          onChange={(e) => updateForm({ ...form, password: e.target.value })}
        />
      </div>
      <button type="submit" className="submit-btn" disabled={loading}>
        Sign Up
      </button>
    </form>
  )
}

function App() {
  const [mode, setMode] = useState<'signin' | 'signup'>('signin')
  const [loggedIn, setLoggedIn] = useState(false)
  const [currentUser, setCurrentUser] = useState<User | null>(null)

  const onLogout = async () => {
    await signOut(auth)
    console.log('signed out')
    setLoggedIn(false)
  }

  useEffect(() => {
    onAuthStateChanged(auth, (user) => {
      if (user) {
        setLoggedIn(true)
        setCurrentUser(user)
      }
    })
  }, [])

  const onGoogleSign = async () => {
    const credentials = await signInWithPopup(auth, googleAuthProvider).catch(
      (error) => {
        console.log(error)
      }
    )
    // send user data to custom backend api
    const user = {
      uid: credentials?.user?.uid,
      name: credentials?.user?.displayName,
      screen_name: credentials?.user?.displayName,
      email: credentials?.user?.email,
      birth_date: '2003-08-30',
    }

    const token = await credentials?.user?.getIdToken()

    const response = await fetch('http://localhost:9000/sign-in', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(user),
    })

    console.log(response)
  }

  const onGithubSign = async () => {
    const credentials = await signInWithPopup(auth, githubAuthProvider).catch(
      (error) => {
        console.log(error)
      }
    )
    // send user data to custom backend api
    const user = {
      uid: credentials?.user?.uid,
      email: credentials?.user?.email,
    }

    const token = await credentials?.user?.getIdToken()

    const response = await fetch('http://localhost:9000/sign-in', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(user),
    })

    console.log(response)
  }

  return (
    <>
      <h1>Auth Demo</h1>
      <p>{loggedIn ? 'Logged in' : 'Not Logged in'}</p>
      {!loggedIn && mode === 'signin' && (
        <a
          href="#"
          onClick={(e) => {
            e.preventDefault()
            setMode('signup')
          }}
        >
          Sign Up
        </a>
      )}
      {!loggedIn && mode === 'signup' && (
        <a
          href="#"
          onClick={(e) => {
            e.preventDefault()
            setMode('signin')
          }}
        >
          Sign In
        </a>
      )}
      {loggedIn && (
        <>
          <div>
            {currentUser?.photoURL && (
              <img
                className="profile-photo"
                src={currentUser.photoURL}
                alt=""
              />
            )}
            <p>UID: {currentUser?.uid}</p>
            <p>Email: {currentUser?.email}</p>
          </div>
          <button onClick={onLogout}>Logout</button>
        </>
      )}
      {!loggedIn && mode === 'signup' && <SignUpForm />}
      {!loggedIn && mode === 'signin' && <SignInForm />}
      {!loggedIn && (
        <div className="social-login">
          <button onClick={onGoogleSign}>Google</button>
          <button onClick={onGithubSign}>GitHub</button>
        </div>
      )}
    </>
  )
}

export default App
