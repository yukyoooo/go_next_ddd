import { ApiContext } from '../../types'
import { fetcher } from '../../utils'

export type GetEmployeeParams = {
  id: number
}

export const getEmployee = async (context: ApiContext, { id }: GetEmployeeParams) => {
  const res = await fetcher(`${context.apiRootUrl.replace(/\/$/g, '')}/employee/${id}`)
  return res
}
