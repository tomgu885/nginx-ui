<script setup lang="ts">
import StdCurd from '@/components/StdDataDisplay/StdCurd.vue'
import cdn_nodes from "@/api/cdn_nodes";
import {antSwitch, input} from "@/components/StdDataEntry";

const columns = [{
    title: () => {return '名称'},
    dataIndex: 'title',
    edit: {
        type: input
    },
},{
    title: () => {return 'api地址'},
    dataIndex: 'api_url',
    edit: {
        type: input
    },
}, {
    title: () => {return '状态'},
    dataIndex: 'state',
    edit: {
        type: antSwitch,
        checkedValue: 1,
        unCheckedValue: 2,

    },
    customRender: (row) => {
        // console.log('row', row.record);
        if (row.record.state == 1) {
            return '启用'
        } else {
            return '停用'
        }
        // return 'state1';
    }
},{
    title:() => {return '操作';},
    dataIndex:'action',
}]

</script>

<template>
<std-curd title="节点" :columns="columns" :api="cdn_nodes"></std-curd>
</template>

<style scoped lang="less">

</style>
