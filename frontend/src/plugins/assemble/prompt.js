import marked from "@plugins/markdown/marked.js";
import "highlight.js/styles/default.css";

const stripHTML = (text) => {
  return text.replace(/<[^>]*>/g, '').replace(/[\r\n\t]/g, '');
}

// assemble frontend chat prompt show
const assembleShowChatList = (promptList, prompt) => {
  let showPromptList = [];
  promptList.forEach((item) => {
    if (item.role !== 'system' && item.content !== '') {
      switch (prompt.type) {
        case 1:
          showPromptList.push({
            role: item.role,
            content: marked.parse(item.content),
          });
          break;
        case 2:
          showPromptList.push({
            role: item.role,
            content: marked.parse(item.prefix + item.content),
          });
          break;
      }
    }
  })
  return showPromptList
}

// assemble request chat list
const assembleReqChatList = (promptList, prompt) => {
  let reqPromptList = [];
  promptList.forEach((item) => {
    if (item.content !== '') {
      switch (prompt.type) {
        case 1:
          reqPromptList.push({
            role: item.role,
            content: stripHTML(item.content),
          });
          break;
        case 2:
          reqPromptList.push({
            role: item.role,
            content: item.prefix + item.content,
          });
          break;
      }
    }
  })
  return reqPromptList;
}

export {
  stripHTML,
  assembleShowChatList,
  assembleReqChatList
}