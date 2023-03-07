'use client'
import useSWR from 'swr'

const IRRemotePage = () => {
  useSWR('/api/ir-remote', () => {
    return null
  })
  return <></>
}

export default IRRemotePage

IRRemotePage.displayName = 'IRRemotePage'
