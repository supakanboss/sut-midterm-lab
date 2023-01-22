import { PrefixInterface } from './IPrefix'
import { InstituteInterface } from './IInstitute'
import { AdminInterface } from './IAdmin'

export interface BranchInterface {
    ID?: number
    Branch_Name: string,
    Branch_Teacher: string,
    Branch_Info: string,

    Prefix: PrefixInterface,
    Institute: InstituteInterface,
    Admin: AdminInterface,
}