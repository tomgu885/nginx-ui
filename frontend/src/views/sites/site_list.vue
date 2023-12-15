<script setup lang="tsx">

import {useGettext} from "vue3-gettext";

import {Badge, Tag} from "ant-design-vue";
import {h, provide} from "vue";
import dayjs from 'dayjs'
import {state2str} from '@/lib/helper/index'
import {Column} from '@/components/StdDesign/types'
import {textarea} from "@/components/StdDataEntry";
import {customRender} from '@/components/StdDesign/StdDataDisplay/StdTableTransformer'
import sites from '@/api/sites'

const {$gettext} = useGettext()

const columns :Column[] = [{
    title: () => { return '域名'},
    dataIndex:'domains',
    sortable: true,
    pithy: true,
},{
    title: '状态',
    dataIndex:'state',
    customRender: (args: customRender) => {
        state2str(args.record.state)
    }
}]


</script>
<template>
    <ACard :title="$gettext('站点')">
        <template #extra>
            <a @click="$router.push('/sites/add')">添加</a>
        </template>
    </ACard>
</template>
<style lang="less">

</style>
