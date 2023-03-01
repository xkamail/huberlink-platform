import React from 'react'

const Form = ({ children }: { children: React.ReactNode }) => {
  return <form className="space-y-4">{children}</form>
}

export default Form

Form.displayName = 'Form'
