'use client'
import AuthService from '@/services/AuthService'
import { useState } from 'react'
import { IUser } from '../types'
import { createProvider } from './index'

export type IUserActions =
  | {
      type: 'fetch-user'
    }
  | { type: 'open-dialog' }
  | { type: 'close-dialog' }
  | { type: 'logout' }

export const [UserContextProvider, useUserDispatch, useUserSelector] =
  createProvider((props: { profile: IUser | null }) => {
    const [userData, setUserData] = useState<IUser | null>(props.profile)

    const dispatch = async (action: IUserActions) => {
      switch (action.type) {
        case 'logout':
          setUserData(null)
          return

          return
        case 'fetch-user':
          await AuthService.me()
            .then((r) => {
              if (!r.success) {
                setUserData(null)
                return
              }
              setUserData(r.data)
            })
            .catch((err) => {
              console.log(err)
              setUserData(null)
            })
          return
      }
    }
    const isLoggedIn = !!userData
    const state = {
      profile: userData,
      isLoggedIn,
    }
    return [state, dispatch]
  })
