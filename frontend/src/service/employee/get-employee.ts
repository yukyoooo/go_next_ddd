import { fetcher } from '@/utils'
import { ApiContext } from '@/types'

export type GetEmployeeParams = {
  id: number
}

const getEmployee = async (context: ApiContext, { id }: GetEmployeeParams) => {
  const res = await fetcher(`${context.apiRootUrl.replace(/\/$/g, '')}/api/employee/${id}`)
  return res
}

export default getEmployee
