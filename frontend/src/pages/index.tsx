import { GetServerSideProps, NextPage } from 'next'
import Head from 'next/head'
import styled from 'styled-components'
import { getEmployee } from '../services/employee/get-employee'
import { ApiContext } from '../types'
import Button from '../components/atoms/Button'

const H1 = styled.h1`
  color: red;
`

const Title = styled.h1`
  font-size: ${({ theme }) => theme.fontSizes[4]};
  color: ${({ theme }) => theme.colors.primary};
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
    <div>
      <Head>
        <title>Create Next App</title>
        <meta name='description' content='Generated by create next app' />
        <link rel='icon' href='/favicon.ico' />
      </Head>

      <main>
        <Title>What&apos;s on your agenda for today?</Title>
        <div>
          Name is: {employee.Name.FirstName} {employee.Name.LastName}.
        </div>
        <Button>Todo</Button>
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
