import {LOCALE_OPTIONS} from "@/locale";
import i18n from '../locale/index'

const {t} = i18n.global

const SysInputEnums = [
  {
    id: 1,
    label: t('common.prompt.select.title1'),
    desc: t('common.prompt.select.title1.desc'),
    type: 2, // 1: chat  2：answer
    is_sys: 2, // 1: is system
    extra: []
  }, {
    id: 2,
    label: t('common.prompt.select.title2'),
    desc: t('common.prompt.select.title2.desc'),
    type: 1,
    is_sys: 1,
    extra: []
  }, {
    id: 3,
    label: t('common.prompt.select.title3'),
    desc: t('common.prompt.select.title3.desc'),
    type: 2,
    is_sys: 2,
    extra: LOCALE_OPTIONS
  }, {
    id: 4,
    label: t('common.prompt.select.title4'),
    desc: t('common.prompt.select.title4.desc'),
    type: 2,
    is_sys: 2,
    extra: []
  }, {
    id: 5,
    label: t('common.prompt.select.title5'),
    desc: t('common.prompt.select.title5.desc'),
    type: 2,
    is_sys: 2,
    extra: []
  }, {
    id: 6,
    label: t('common.prompt.select.title6'),
    desc: t('common.prompt.select.title6.desc'),
    type: 2,
    is_sys: 2,
    extra: []
  }, {
    id: 7,
    label: t('common.prompt.select.title7'),
    desc: t('common.prompt.select.title7.desc'),
    type: 2,
    is_sys: 2,
    extra: []
  }, {
    id: 8,
    label: t('common.prompt.select.title8'),
    desc: t('common.prompt.select.title8.desc'),
    type: 2,
    is_sys: 2,
    extra: []
  },
];

export default SysInputEnums