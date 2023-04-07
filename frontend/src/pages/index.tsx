import Head from 'next/head'
import { ApiContext } from '../types'
import { GetServerSideProps, NextPage } from 'next'
import { getEmployee } from '@/service/employee/get-employee'
import styles from '../styles/Home.module.css'
import styled from 'styled-components'

const H1 = styled.h1`
  color: red;
`

type SSRProps = {
  employee: {
    ID: number
    Name: {
      FirstName: string
      LastName: string
    }
    Email: { Value: string }
    Password: { Value: string }
    Role: number
  }
}

const SSR: NextPage<SSRProps> = ({ employee }: SSRProps) => {
  if (!employee) return <div>loading...</div>

  return (
    <div className={styles.container}>
      <Head>
        <title>Create Next App</title>
        <meta name='description' content='Generated by create next app' />
        <link rel='icon' href='/favicon.ico' />
      </Head>

      <main className={styles.main}>
        <H1 className={styles.title}>
          Welcome to <a href='https://nextjs.org'>Next.js!</a>
        </H1>
        <div>
          Name is: {employee.Name.FirstName} {employee.Name.LastName}.
        </div>
      </main>
    </div>
  )
}

export const getServerSideProps: GetServerSideProps<SSRProps> = async () => {
  const context: ApiContext = {
    apiRootUrl: process.env.API_BASE_URL || 'http://localhost:8090',
  }
  const employee = await getEmployee(context, { id: 1 })
  return {
    props: {
      employee,
    },
  }
}

export default SSR
