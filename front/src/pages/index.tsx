import type { NextPage } from 'next'
import { useEffect, useRef, useState } from 'react'

const Home: NextPage = () => {
  const socketRef = useRef<WebSocket>()
  const [isConnected, setIsConnected] = useState(false)
  const [formMessage, setFormMessage] = useState('')
  const [sentMessage, setSentMessage] = useState('')

  const sendData = (event: any) => {
    event.preventDefault()
    setFormMessage(event.target[0].value)
    socketRef.current?.send(event.target[0].value)
  }

  useEffect(() => {
    socketRef.current = new WebSocket('ws://' + process.env.NEXT_PUBLIC_BACKEND_SERVER + process.env.NEXT_PUBLIC_PORT + '/socket')
    socketRef.current.onopen = function () {
      setIsConnected(true)
      console.log('Connected')
    }

    socketRef.current.onclose = function () {
      console.log('closed')
      setIsConnected(false)
    }

    // server 側から送られてきたデータを受け取る
    socketRef.current.onmessage = function (event) {
      setSentMessage(event.data)
    }

    return () => {
      if (socketRef.current == null) {
        return
      }
      socketRef.current.close()
    }
  }, [])

  return (
    <>
      <h1>WebSocket is connected : {`${isConnected}`}</h1>
      <form onSubmit={sendData}>
        <input type='text' name='socketData' />
        <button type='submit'>Server に送信</button>
      </form>
      <h3>form message: {formMessage}</h3>
      <h3>sent message: {sentMessage}</h3>
    </>
  )
}

export default Home
