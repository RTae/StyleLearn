<template>
  <div id="app">
    <v-app>
      <Header/>
      <Content/>
      <Footer/>
    </v-app>
  </div>
</template>

<script>
import Header from '@/components/core/Header'
import Content from '@/components/core/Content'
import Footer from '@/components/core/Footer'

export default {
  name: 'App',
  components: {
    Header,
    Content,
    Footer
  }
}
</script>
