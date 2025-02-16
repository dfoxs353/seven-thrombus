import { useState } from "react"

export const useVisible = (initial: boolean) => {
  const [visible, set] = useState<boolean>(initial)

  const toggle = () => {
    set(prev => !prev)
  }

  const close = () => {
    set(false)
  }

  const open = () => {
    set(true)
  }

  return {
    visible,
    toggle,
    close,
    open
  }

}