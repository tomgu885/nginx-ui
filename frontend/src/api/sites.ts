import Crud from '@/api/curd'
import http from '@/lib/http'

import {AxiosRequestConfig} from "axios/index";

class Sites extends Crud {
    protected readonly plural: string = 'sites'
    reloadAllSites(config: AxiosRequestConfig) {
        return http.post(this.baseUrl+'/reload')
    }
    enable(id: number, config : AxiosRequestConfig) {
        return http.post(this.baseUrl+'/' +id+'/enable')
    }

    disable(id: number, config: AxiosRequestConfig ) {

    }
}

const sites = new Sites('/sites')
export default sites
