import {marked} from 'marked';
import hljs from "highlight.js";

marked.setOptions({
  renderer: new marked.Renderer(),
  highlight: function (code, lang) {
    let language = hljs.getLanguage(lang) ? lang : 'plaintext';
    if (lang === 'vue') {
      language = 'xml';
    }
    return hljs.highlight(code, {language},true).value;
  },
  langPrefix: 'hljs language-', // highlight.js css expects a top-level 'hljs' class.
  pedantic: false,
  gfm: true,
  breaks: false,
  sanitize: false,
  smartLists: true,
  smartypants: false,
  xhtml: false,
});

export default marked