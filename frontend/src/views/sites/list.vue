<script setup lang="ts">
import StdTable from '@/components/StdDataDisplay/StdTable.vue'
import {customRender, datetime} from '@/components/StdDataDisplay/StdTableTransformer'
import {useGettext} from 'vue3-gettext'
import {state2str} from '@/lib/helper/index'
import sites from '@/api/sites'
import {Badge, message} from 'ant-design-vue'
import {h, ref} from 'vue';

import StdCurd from '@/components/StdDataDisplay/StdCurd.vue'
import {input, textarea, antSwitch, selector, select, radio} from '@/components/StdDataEntry'

import SiteDuplicate from "@/views/domain/components/SiteDuplicate.vue";
import {stateFormat} from "@/lib/helper";

const {$gettext, interpolate} = useGettext()

const columns =[
// {
//     title: () => {return '名称'},
//     help: () => {return 'help内容'},
//     dataIndex:'name',
//     sorter: false,
//     show: true,
//     pithy: true,
//     edit: {
//         // placeholder: () => {return '说明'},
//         // type: input,
//     },
//     search: true,
// },
    {
    title:  '域名',
    dataIndex:'domains',
    edit: {
        placeholder:  '多个域名以空格 或者换行隔开',
        type: textarea,
        rows: 3,
    },
    // search: true,
},{
    title:'源站',
        dataIndex:'upstream_ips',
        edit: {
            type:textarea,
            placeholder:'多个以换行分隔',
            rows: 4,
        }
    },{
    title: () => {return '回源策略'},
    customRender: (args) => { // 1:同端口协议, 2: 回落到 80, 3: 回落到 443
        console.log('args', args.record.upstream_port_policy)
    },
    dataIndex:'upstream_port_policy',
    edit: {
        type: radio,
        mask:
            {
                1: '同端口协议',
                2: '回落到http(80)',
                3:'回落到https(443)'
            }
        ,
        disable_search: true,
    }
},{
    title:'回源host',
    dataIndex:'upstream_host',
},{
    title:'源站',
        dataIndex:'upstream_ips',

    }
,{
    title: () => 'http端口',
    dataIndex:'http_ports',
    edit: {
        placeholder: () => { return '多个端口以空格隔开'},
        type: input,

    }
},
{
    title: () => '启用https',
    dataIndex:'ssl_enable',
  customRender: (args) => {
      if (args.record.ssl_enable != 1 ) {
        return '否';
      } else {
        // ssl_cert_state	 ssl 证书 状态 1:等待, 2: 申请开始, 3: 已完成, 4:失败
        switch (args.record.ssl_cert_state) {
          case 1:
            return '等待申请'
          case 2:
            return '申请中'
          case 3:
            return '正常'
          case 4:
            return '申请失败'
        }
        return '未知'
      }
  },
    edit: {
        type: antSwitch,
        checkedValue: 1,
        unCheckedValue: 2,
    }
},{
  title: () => 'https端口',
  dataIndex:'https_ports',
    edit: {
        placeholder: () => {return '多个端口以空格隔开'},
      type: input,
    },
},{
    title:() => 'websocket',
    dataIndex:'websocket_enable',
    customRender: (args) => {
        return state2str(args.record.websocket_enable)
    },
    edit: {
        type: antSwitch,
        checkedValue: 1,
        unCheckedValue: 2,
    }
},
{
    title: () => {return '跳转链接'},
    dataIndex:'redirect',
    edit:{
        type: input,
        placeholder: () => {return '当该项功能开启时，其它设置将失效'}
    }
},
{
  title: '操作',
  dataIndex:'action'
}
]

const table = ref(null)

interface Table {
    get_list() : void
}

function destroy(siteId: number) {
    sites.destroy(siteId).then(() => {
        const t: Table | null = table.value
        t!.get_list()
        message.success(interpolate($gettext('Delete site: %{siteId}'), {site_id:siteId}))
    }).catch((e:any) => {
        message.error(e?.message?? $gettext('Server error'))
    })
}

const show_duplicator = ref(false)
const target = ref('')

function handle_click_duplicate(name: string) {
    show_duplicator.value = true
    target.value = name
}

function add(item) {
  console.log('add', item);
}

const selectedRowKeys = ref([])
const props = defineProps({
    api:Object,
    columns: Array,
    title: String,
})
function edit() {

}


</script>

<template>
    <std-curd title="CDN加速站点" :columns="columns" :api="sites" di1sable_add="true" />
</template>

<style scoped lang="less">

</style>
