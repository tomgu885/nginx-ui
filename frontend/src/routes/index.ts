import {createRouter, createWebHashHistory} from 'vue-router'
import gettext from '../gettext'
import {useUserStore} from '@/pinia'

import {
    CloudOutlined,
    CodeOutlined,
    DatabaseOutlined,
    FileOutlined,
    FileTextOutlined,
    HomeOutlined,
    InfoCircleOutlined,
    SafetyCertificateOutlined,
    SettingOutlined,
    UserOutlined
} from '@ant-design/icons-vue'
import NProgress from 'nprogress'
import 'nprogress/nprogress.css'

const {$gettext} = gettext

export const routes = [
    {
        path: '/',
        name: () => '首页',
        component: () => import('@/layouts/BaseLayout.vue'),
        redirect: '/dashboard',
        children: [
            {
                path: 'dashboard',
                component: () => import('@/views/dashboard/DashBoard.vue'),
                name: () => '仪表板',
                meta: {
                    // hiddenHeaderContent: true,
                    icon: HomeOutlined
                },
                redirect: '/sites/list',

            },
            {
                path:"sites",
                name :() => '网站管理',
                component:() => import('@/views/sites/SiteList.vue'),
                meta: {
                    icon: CloudOutlined,
                },
                children:[{
                    path: 'add',
                    name:() => '添加',
                    component: ()=> import('@/view/sites/SiteAdd.vue'),
                    meta: {
                        hiddenInSidebar: true
                    }
                }]
            },
            {
                path: 'domain',
                name: () => $gettext('Manage Sites'),
                component: () => import('@/layouts/BaseRouterView.vue'),
                meta: {
                    icon: CloudOutlined
                },
                redirect: '/domain/list',
                children: [{
                    path: 'list',
                    name: () => $gettext('Sites List'),
                    component: () => import('@/views/domain/DomainList.vue')
                }, {
                    path: 'add',
                    name: () => $gettext('Add Site'),
                    component: () => import('@/views/domain/DomainAdd.vue')
                }, {
                    path: ':name',
                    name: () => $gettext('Edit Site'),
                    component: () => import('@/views/domain/DomainEdit.vue'),
                    meta: {
                        hiddenInSidebar: true
                    }
                }]
            },
            {
                path: 'config',
                name: () => $gettext('Manage Configs'),
                component: () => import('@/views/config/Config.vue'),
                meta: {
                    hiddenInSidebar: true,
                    icon: FileOutlined,
                    hideChildren: true
                }
            },
            {
                path: 'config/:name+/edit',
                name: () => $gettext('Edit Configuration'),
                component: () => import('@/views/config/ConfigEdit.vue'),
                meta: {
                    hiddenInSidebar: true
                }
            },
            {
                path: 'cert',
                name: () => $gettext('Certification'),
                component: () => import('@/layouts/BaseRouterView.vue'),
                meta: {
                    icon: SafetyCertificateOutlined
                },
                children: [
                    {
                        path: 'list',
                        name: () => $gettext('Certification List'),
                        component: () => import('@/views/cert/Cert.vue')
                    },
                    {
                        path: 'dns_credential',
                        name: () => $gettext('DNS Credentials'),
                        component: () => import('@/views/cert/DNSCredential.vue')
                    }
                ]
            },

            {
                path: 'environment',
                name: () => $gettext('Environment'),
                component: () => import('@/views/environment/Environment.vue'),
                meta: {
                    icon: DatabaseOutlined
                }
            },
            // {
            //     path: 'user',
            //     name: () => $gettext('Manage Users'),
            //     component: () => import('@/views/user/User.vue'),
            //     meta: {
            //         icon: UserOutlined
            //     }
            // },
            {
                path: 'preference',
                name: () => $gettext('Preference'),
                component: () => import('@/views/preference/Preference.vue'),
                meta: {
                    icon: SettingOutlined
                }
            },
            {
                path: 'system',
                name: () => $gettext('System'),
                redirect: 'system/about',
                meta: {
                    icon: InfoCircleOutlined
                },
                children: [{
                    path: 'about',
                    name: () => $gettext('About'),
                    component: () => import('@/views/system/About.vue')
                }, {
                    path: 'upgrade',
                    name: () => $gettext('Upgrade'),
                    component: () => import('@/views/system/Upgrade.vue')
                }]
            }
        ]
    },
    {
        path: '/login',
        name: () => $gettext('Login'),
        component: () => import('@/views/other/Login.vue'),
        meta: {noAuth: true}
    },
    {
        path: '/:pathMatch(.*)*',
        name: () => $gettext('Not Found'),
        component: () => import('@/views/other/Error.vue'),
        meta: {noAuth: true, status_code: 404, error: () => $gettext('Not Found')}
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    // @ts-ignore
    routes: routes
})

NProgress.configure({showSpinner: false})

router.beforeEach((to, from, next) => {
    // @ts-ignore
    document.title = to.name?.() + ' | Nginx UI'

    NProgress.start()

    const user = useUserStore()
    const {is_login} = user

    if (to.meta.noAuth || is_login) {
        next()
    } else {
        next({path: '/login', query: {next: to.fullPath}})
    }

})

router.afterEach(() => {
    NProgress.done()
})

export default router
