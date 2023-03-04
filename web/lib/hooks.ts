import { useUserDispatch, useUserSelector } from './contexts/UserContext'

export {}

export const useUser = () => {
  const dispatch = useUserDispatch()
  const isLoggedIn = useUserSelector((s) => s.isLoggedIn)
  const isPrepare = useUserSelector(
    (s) => s.status === 'idle' || s.status === 'loading'
  )
  const profile = useUserSelector((s) => s.profile)

  const logout = () => {
    dispatch({ type: 'logout' })
  }

  return {
    profile: profile!,
    isLoggedIn,
    logout,
    isPrepare,
  }
}
