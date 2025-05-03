const svg_messenger = `
<svg xmlns="http://www.w3.org/2000/svg" height="0.9em" fill="currentColor" viewBox="0 0 16 16" style="vertical-align: middle; margin: 0;">
    <path d="M0 7.76C0 3.301 3.493 0 8 0s8 3.301 8 7.76-3.493 7.76-8 7.76c-.81 0-1.586-.107-2.316-.307a.64.64 0 0 0-.427.03l-1.588.702a.64.64 0 0 1-.898-.566l-.044-1.423a.64.64 0 0 0-.215-.456C.956 12.108 0 10.092 0 7.76m5.546-1.459-2.35 3.728c-.225.358.214.761.551.506l2.525-1.916a.48.48 0 0 1 .578-.002l1.869 1.402a1.2 1.2 0 0 0 1.735-.32l2.35-3.728c.226-.358-.214-.761-.551-.506L9.728 7.381a.48.48 0 0 1-.578.002L7.281 5.98a1.2 1.2 0 0 0-1.735.32z"/>
</svg>`
let two_bubbles =`
<svg height=" 1.2em" style="vertical-align: middle; margin: 0;" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg" preserveAspectRatio="xMinYMin"
    class="jam jam-messages-alt-f">
    <path
        d="M7.46 2.332C8.74.913 10.746 0 13 0c3.866 0 7 2.686 7 6 0 1.989-1.13 3.752-2.868 4.844L17 10.92v4.05a1 1 0 0 1-1.718.696l-1.14-1.174c1.069-1.264 1.698-2.816 1.698-4.493 0-4.067-3.698-7.395-8.38-7.667" />
    <path
        d="m8.385 15.886-3.667 3.78A1 1 0 0 1 3 18.97v-4.05l-.132-.076C1.129 13.752 0 11.989 0 10c0-3.314 3.134-6 7-6s7 2.686 7 6c0 2.726-2.121 5.028-5.026 5.758a8 8 0 0 1-.589.128" />
</svg>`


const svg_send = `
<svg height="0.9em" style="vertical-align: middle; margin: 0;" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
    <path
        d="m20.34 9.32-14-7a3 3 0 0 0-4.08 3.9l2.4 5.37a1.06 1.06 0 0 1 0 .82l-2.4 5.37A3 3 0 0 0 5 22a3.14 3.14 0 0 0 1.35-.32l14-7a3 3 0 0 0 0-5.36Zm-.89 3.57-14 7a1 1 0 0 1-1.35-1.3l2.39-5.37a2 2 0 0 0 .08-.22h6.89a1 1 0 0 0 0-2H6.57a2 2 0 0 0-.08-.22L4.1 5.41a1 1 0 0 1 1.35-1.3l14 7a1 1 0 0 1 0 1.78" />
</svg>`

const svg_logout = `    
<svg height="0.9em" style="vertical-align: middle; margin: 0;" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" fill="none">
    <path stroke="#000000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12h-9.5m7.5 3l3-3-3-3m-5-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2h5a2 2 0 002-2v-1"/>
</svg>`

const svg_post_creation =`
<svg  height="1.2em" viewBox="0 0 32 32" style="
    fill-rule: evenodd;
    clip-rule: evenodd;
    stroke-linejoin: round;
    stroke-miterlimit: 2;
  " version="1.1" xml:space="preserve" xmlns="http://www.w3.org/2000/svg" xmlns:serif="http://www.serif.com/"
    xmlns:xlink="http://www.w3.org/1999/xlink">
    <g id="SVGRepo_bgCarrier" stroke-width="0"></g>
    <g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
    <g id="SVGRepo_iconCarrier">
        <g id="Layer1">
            <path
                d="M16,6l-13,0c-0.552,0 -1,0.448 -1,1l0,22c0,0.552 0.448,1 1,1l22,0c0.552,0 1,-0.448 1,-1l0,-13c0,-0.552 -0.448,-1 -1,-1c-0.552,-0 -1,0.448 -1,1l0,12c0,0 -20,0 -20,0c0,0 0,-20 0,-20c-0,0 12,0 12,0c0.552,0 1,-0.448 1,-1c0,-0.552 -0.448,-1 -1,-1Zm-9,19l14,-0c0.552,0 1,-0.448 1,-1c0,-0.552 -0.448,-1 -1,-1l-14,0c-0.552,0 -1,0.448 -1,1c0,0.552 0.448,1 1,1Zm-0,-4l4,0c0.552,-0 1,-0.448 1,-1c-0,-0.552 -0.448,-1 -1,-1l-4,0c-0.552,-0 -1,0.448 -1,1c-0,0.552 0.448,1 1,1Zm22.707,-13.293c0.391,-0.39 0.391,-1.024 0,-1.414l-4,-4c-0.39,-0.391 -1.024,-0.391 -1.414,-0l-10,10c-0.14,0.139 -0.235,0.317 -0.274,0.511l-1,5c-0.065,0.328 0.037,0.667 0.274,0.903c0.236,0.237 0.575,0.339 0.903,0.274l5,-1c0.194,-0.039 0.372,-0.134 0.511,-0.274l10,-10Zm-22.707,9.293l4,0c0.552,0 1,-0.448 1,-1c0,-0.552 -0.448,-1 -1,-1l-4,0c-0.552,0 -1,0.448 -1,1c0,0.552 0.448,1 1,1Zm0,-4l5,-0c0.552,0 1,-0.448 1,-1c0,-0.552 -0.448,-1 -1,-1l-5,-0c-0.552,0 -1,0.448 -1,1c0,0.552 0.448,1 1,1Z">
            </path>
        </g>
    </g>
</svg>`
export default { svg_messenger, two_bubbles, svg_send, svg_logout, svg_post_creation };