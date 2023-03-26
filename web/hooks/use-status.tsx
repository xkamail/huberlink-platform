import { useState } from 'react'

export function useStatus(
  initial: 'idle' | 'loading' | 'ok' | 'error' | 'notfound' = 'idle'
) {
  return useState<'idle' | 'loading' | 'ok' | 'error' | 'notfound'>(initial)
}
