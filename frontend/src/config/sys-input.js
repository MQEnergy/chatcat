import {LOCALE_OPTIONS} from "@/locale";

const SysInputEnums = [
  {
    id: 1,
    label: '对话:',
    desc: "",
    type: 1, // 1: 连续对话  2：问答
    is_sys: 1, // 1: 是否需要system
    extra: []
  }, {
    id: 2,
    label: '问答:',
    desc: "请输入问答内容:",
    type: 2,
    is_sys: 2,
    extra: []
  }, {
    id: 3,
    label: '翻译:',
    desc: "请将以下翻译成",
    type: 2,
    is_sys: 2,
    extra: LOCALE_OPTIONS
  }, {
    id: 4,
    label: '改写:',
    desc: "重写以下内容:",
    type: 2,
    is_sys: 2,
    extra: []
  }, {
    id: 5,
    label: '润色:',
    desc: "将以下内容进行润色:",
    type: 2,
    is_sys: 2,
    extra: []
  }, {
    id: 6,
    label: '总结:',
    desc: "将一下内容进行归纳总结:",
    type: 2,
    is_sys: 2,
    extra: []
  }, {
    id: 7,
    label: '解释:',
    desc: "解释以下内容:",
    type: 2,
    is_sys: 2,
    extra: []
  }, {
    id: 8,
    label: '解释代码:',
    desc: "解释以下代码:",
    type: 2,
    is_sys: 2,
    extra: []
  },
];

export default SysInputEnums