<template>
  <router-view/>
</template>
<script setup>
import {GetGeneralInfo} from "../wailsjs/go/setting/Service.js";

if (window.go !== undefined) {
  GetGeneralInfo().then(res => {
    if (res.code === 0) {
      const theme = res.data.theme;
      switch (theme) {
        case 1:
          const darkThemeMq = window.matchMedia("(prefers-color-scheme: dark)");
          darkThemeMq.addListener(e => {
            console.log(e)
            if (e.matches) {
              document.body.setAttribute('arco-theme', 'dark');
            } else {
              document.body.removeAttribute('arco-theme');
            }
          });
        case 2:
          document.body.removeAttribute('arco-theme');
          break;
        case 3:
          document.body.setAttribute('arco-theme', 'dark');
          break;
      }
    }
  })
}
</script>

<style>

</style>
