// APIコンテキスト
export type ApiContext = {
  apiRootUrl: string
}

export type Employee = {
  ID: number
  Name: {
    FirstName: string
    LastName: string
  }
  Email: { Value: string }
  Password: { Value: string }
  Role: number
}
