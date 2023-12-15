import StdDataEntry from './StdDataEntry.js'
import {h} from 'vue'
import {Input, InputNumber, Textarea, Switch} from 'ant-design-vue'
import StdSelector from './components/StdSelector.vue'
import StdSelect from './components/StdSelect.vue'
import StdRadio from './components/StdRadio.vue'
import StdPassword from './components/StdPassword.vue'
import placeholder from "lodash/fp/placeholder";
import type { StdDesignEdit } from '@/components/StdDesign/types'

interface IEdit {
    type: Function
    placeholder: any
    mask: any
    key: any
    value: any
    recordValueIndex: any
    selectionType: any
    api: Object,
    columns: any,
    data_key: any,
    disable_search: boolean,
    get_params: Object,
    description: string
    generate: boolean
    min: number
    max: number,
    extra: string
}

function fn(obj: Object, desc: any) {
    let arr: string[]
    if (typeof desc === 'string') {
        arr = desc.split('.')
    } else {
        arr = [...desc]
    }

    while (arr.length) {
        // @ts-ignore
        const top = obj[arr.shift()]
        if (top === undefined) {
            return null
        }
        obj = top
    }
    return obj
}

function readonly(edit: IEdit, dataSource: any, dataIndex: any) {
    return h('p', fn(dataSource, dataIndex))
}

function placeholder_helper(edit: StdDesignEdit) {
    return typeof edit.config?.placeholder === 'function' ? edit.config?.placeholder() : edit.config?.placeholder
}

function input(edit: StdDesignEdit, dataSource: any, dataIndex: any) {

    return h(Input, {
        placeholder: placeholder_helper(edit),
        value: dataSource?.[dataIndex],
        'onUpdate:value': value => {
            dataSource[dataIndex] = value
        }
    })
}

function inputNumber(edit: StdDesignEdit, dataSource: any, dataIndex: any) {
    return h(InputNumber, {
        placeholder: placeholder_helper(edit),
        min: edit.min,
        max: edit.max,
        value: dataSource?.[dataIndex],
        'onUpdate:value': value => {
            dataSource[dataIndex] = value
        }
    })
}

function textarea(edit: IEdit, dataSource: any, dataIndex: any) {
    let plh = ''
    if ('placeholder' in edit) {
        plh = typeof edit.placeholder == 'string' ? edit.placeholder : edit.placeholder()
    }
    return h(Textarea, {
        placeholder: plh,
        rows: edit.rows? edit.rows:4,
        value: dataSource?.[dataIndex],
        'onUpdate:value': value => {
            dataSource[dataIndex] = value
        }
    })
}

function password(edit: IEdit, dataSource: any, dataIndex: any) {
    return <StdPassword
        v-model:value={dataSource[dataIndex]}
        generate={edit.generate}
        placeholder={edit.placeholder}
    />
}

function radio(edit: IEdit, dataSource: any, dataIndex: any) {
    return <StdRadio
        v-model:value={dataSource[dataIndex]}
        mask={edit.mask}
    />
}

function select(edit: IEdit, dataSource: any, dataIndex: any) {
    return <StdSelect
        v-model:value={dataSource[dataIndex]}
        mask={edit.mask}
    />
}

function selector(edit: IEdit, dataSource: any, dataIndex: any) {
    return <StdSelector
        v-model:selectedKey={dataSource[dataIndex]}
        value={edit.value}
        recordValueIndex={edit.recordValueIndex}
        selectionType={edit.selectionType}
        api={edit.api}
        columns={edit.columns}
        data_key={edit.data_key}
        disable_search={edit.disable_search}
        get_params={edit.get_params}
        description={edit.description}
    />
}

function antSwitch(edit: IEdit, dataSource: any, dataIndex: any) {
    // console.log('antSwitch Edit', edit.checkedValue);
    return h(Switch, {
        checkedValue: edit.checkedValue?? true,
        unCheckedValue: edit.unCheckedValue?? true,
        checked: dataSource?.[dataIndex],
        'onUpdate:checked': value => {
            dataSource[dataIndex] = value
        }
    })
}

export {
    readonly,
    input,
    textarea,
    radio,
    select,
    selector,
    password,
    inputNumber,
    antSwitch
}

export default StdDataEntry
