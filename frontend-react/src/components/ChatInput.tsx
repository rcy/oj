import { ChangeEventHandler, KeyboardEventHandler, useState } from "react"

interface Props {
  onSubmit: Function
}

export default function ChatInput({ onSubmit }: Props) {
  const [input, setInput] = useState('')

  const handleChange: ChangeEventHandler<HTMLInputElement> = (ev) => {
    setInput(ev.target.value)
  }

  const handleKey: KeyboardEventHandler<HTMLInputElement> = (ev) => {
    if (ev.key === 'Enter') {
      if (onSubmit(input)) {
        setInput('')
      }
    }
  }

  return (
    <input
      className='border border-solid border-gray-300'
      type="text"
      placeholder="say something"
      onChange={handleChange}
      onKeyDown={handleKey}
      value={input}
    />
  )
}
