import {LOCALE_OPTIONS} from "@/locale";

const SysInputEnums = [
  {
    label: '对话:',
    value: "",
    type: 0,
    extra: []
  }, {
    label: '问答:',
    value: "请输入问答内容:",
    type: 1,
    extra: []
  }, {
    label: '翻译:',
    value: "请将以下翻译成:",
    type: 2,
    extra: LOCALE_OPTIONS
  }, {
    label: '改写:',
    value: "重写以下内容:",
    type: 3,
    extra: []
  }, {
    label: '润色:',
    value: "将以下内容进行润色:",
    type: 4,
    extra: []
  }, {
    label: '总结:',
    value: "总结:",
    type: 5,
    extra: []
  }, {
    label: '解释:',
    value: "解释以下内容:",
    type: 7,
    extra: []
  }, {
    label: '解释代码:',
    value: "解释以下代码:",
    type: 8,
    extra: []
  },
];

export default SysInputEnums