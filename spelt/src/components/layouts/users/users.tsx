import { getRouteApi } from '@tanstack/react-router'
import { Main } from '@/components/layouts/main'
import { users } from './items'
import { UsersPrimaryButtons } from './users-primary-buttons'
import { UsersProvider } from './users-provider'
import { UsersTable } from './users-table'

const route = getRouteApi('/')

export function Users() {
  const nv = route.useNavigate()
  const search = route.useSearch()

  return (
    <UsersProvider>
      <Main>
        <div className="mb-2 flex flex-wrap items-center justify-between space-y-2">
          <div>
            <h2 className="text-2xl font-bold tracking-tight">User List</h2>
            <p className="text-muted-foreground">
              Manage your users and their roles here.
            </p>
          </div>
          <UsersPrimaryButtons />
        </div>
        <div className="-mx-4 flex-1 overflow-auto px-4 py-1 lg:flex-row lg:space-y-0 lg:space-x-12">
          <UsersTable data={users} search={search} navigate={nv} />
        </div>
      </Main>
    </UsersProvider>
  )
}
