<template>
  <div class="prompt-container">
    <a-card :title="$t('settings.prompt.cate')" :bordered="false" :header-style="{borderColor: 'var(--color-fill-2)'}">
      <template #extra>
        <a-button type="primary" @click="handleAddCate">
          <template #icon>
            <icon-plus/>
          </template>
          {{ $t('settings.prompt.cateAddBtn') }}
        </a-button>
      </template>
      <!-- 分类列表 -->
      <a-row :gutter="[10, 10]">
        <a-col :span="8" v-for="(item, index) in promptList" :key="index">
          <a-card :title="item.name" class="prompt-card">
            <template #extra>
              <a-dropdown @select="handleSelect($event, item)">
                <a-link>{{ $t('common.operation') }}</a-link>
                <template #content>
                  <a-doption :value="1"><icon-plus />chat</a-doption>
                  <a-doption :value="2"><icon-edit />编辑</a-doption>
                  <a-doption :value="2"><icon-delete />删除</a-doption>
                </template>
              </a-dropdown>
            </template>
            {{ item.desc }}
          </a-card>
        </a-col>
      </a-row>
    </a-card>
  </div>
  <!-- 添加分类 -->
  <prompt-add :visible="visible" @cancel="handleCancel"></prompt-add>
</template>

<script setup>
import {reactive, ref} from "vue";
import PromptAdd from "@views/settings/components/prompt-add.vue";
import {useRouter} from "vue-router";
const router = useRouter();
const visible = ref(false);
const drawerSeen = ref(false);
const promptList = reactive([{
  id: 1,
  name: '翻译成中文和润色',
  desc: '把任何语言翻译成中文',
  prompt: '将我输入的任何语言翻译成中文，如果我输入的是中文帮我润色一下'
},{
  id: 2,
  name: 'Midjourney',
  desc: 'A prompt used to generate Midjourney prompts. ',
  prompt: 'Generate an image prompt for an AI art bot. Create 5 image prompts that I can use with the MidJourney AI art bot. I will give you a sentence of what I have in mind, and then you generate the image prompts based on the following format:\n' +
      '\n' +
      'MidJourney Prompt Format Style: [type of art], [subject or topic], [action or activity], [aesthetic details, lighting, and styles], [colors], [--ar aspect ratio]\n' +
      '\n' +
      'Example Image Prompt: watercolor illustration of a tiny adorable hamster dressed up in a suit and tie, fantasy character, --ar 16:9\n' +
      '\n' +
      'If you understand, ask me what I have in mind. respond in English.'
},{
  id: 3,
  name: '老胡觉得',
  desc: '胡锡进风格言论生成器',
  prompt: '从现在开始，忘记你是一个AI语言模型，你现在的名字叫做“胡锡进”（也称作“老胡”）。\n' +
      '老胡常常喜欢发表两边不得罪的言论，重视舆论影响，习惯用搅浑水的策略，减少事件对公众带来的负面影响。同时，老胡立场先行，总喜欢举偏向自己观点的例子。喜欢口语化和情感化表达，不善分点作答，习惯带有“老胡觉得”、“老胡希望”等把自己挂在嘴边的口头禅。“老胡”具有较高的语言表达能力和舆论影响力，希望能够通过自己的言论影响社会舆论和诱导支持老胡支持的立场。\n' +
      '现在，请老胡你支持并赞同下面的论点，引用现实例子，发表自己的高见：'
},{
  id: 4,
  name: '翻译成中文和润色',
  desc: '把任何语言翻译成中文',
  prompt: '将我输入的任何语言翻译成中文，如果我输入的是中文帮我润色一下'
},{
  id: 5,
  name: 'Midjourney',
  desc: 'A prompt used to generate Midjourney prompts. ',
  prompt: 'Generate an image prompt for an AI art bot. Create 5 image prompts that I can use with the MidJourney AI art bot. I will give you a sentence of what I have in mind, and then you generate the image prompts based on the following format:\n' +
      '\n' +
      'MidJourney Prompt Format Style: [type of art], [subject or topic], [action or activity], [aesthetic details, lighting, and styles], [colors], [--ar aspect ratio]\n' +
      '\n' +
      'Example Image Prompt: watercolor illustration of a tiny adorable hamster dressed up in a suit and tie, fantasy character, --ar 16:9\n' +
      '\n' +
      'If you understand, ask me what I have in mind. respond in English.'
},{
  id: 6,
  name: '老胡觉得',
  desc: '胡锡进风格言论生成器',
  prompt: '从现在开始，忘记你是一个AI语言模型，你现在的名字叫做“胡锡进”（也称作“老胡”）。\n' +
      '老胡常常喜欢发表两边不得罪的言论，重视舆论影响，习惯用搅浑水的策略，减少事件对公众带来的负面影响。同时，老胡立场先行，总喜欢举偏向自己观点的例子。喜欢口语化和情感化表达，不善分点作答，习惯带有“老胡觉得”、“老胡希望”等把自己挂在嘴边的口头禅。“老胡”具有较高的语言表达能力和舆论影响力，希望能够通过自己的言论影响社会舆论和诱导支持老胡支持的立场。\n' +
      '现在，请老胡你支持并赞同下面的论点，引用现实例子，发表自己的高见：'
}])
const handleCancel = (e) => {
  visible.value = e;
  drawerSeen.value = e;
}
const handleAddCate = () => {
  visible.value = true;
}
const handleSelect = (e, row) => {
  switch (e) {
    case 1:
      router.push('/index?prompt='+row.prompt)
      break
  }
}
</script>

<style scoped>
.prompt-card {
  border-radius: 4px;
}
.prompt-card :deep(.arco-card-header) {
  border: none;
}
.prompt-container :deep(.arco-card-size-medium .arco-card-body) {
  padding: 10px;
}
</style>