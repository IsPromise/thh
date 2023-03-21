
<script setup>
import MarkdownIt from "markdown-it"
// import Mathjax from "mathjax"
import { ref, onMounted } from 'vue'

const name = ""
const props = {}

onMounted(() => {
  this.$formula(document.getElementById('katexData'))
})


const md = new MarkdownIt();
const p = function (md, option) {
  console.log(md)
  console.log(option)
};
md.use(p)
const result = md.render(`
# markdown-it test!

## two

### three

#### four

##### five


\`\`\`sql
insert into dogs values(1,2,3);
\`\`\`
- [ ] 111
- [ ] 112

[baidu](https://www.baidu.com)

|12|12|
|---|---|
|1 1| 112|

$\\sqrt{3x-1}+(1+x)^2$

    `);
</script>
<template>
  <div v-html="result">
  </div>
  <div class="mermaid">
    sequenceDiagram
    Alice->>+John: Hello John, how are you?
    Alice->>+John: John, can you hear me?
    John-->>-Alice: Hi Alice, I can hear you!
    John-->>-Alice: I feel great!
  </div>
  <p id="katexData">
    $\int$
  </p>
</template>
