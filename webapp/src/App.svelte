<script>
	import Navbar from "./Navbar.svelte";
	import Login from "./login.svelte";
	console.log("Hola mu7ndo")
	function getCookie(cookiename) {
        var cookies = document.cookie.split(';');
		console.log("1  " + document.cookie)
        var ret = '';
        for(var i = 0; i < cookies.length; i++) {
			if (cookies[i].trim().startsWith(cookiename)){
				return cookies[i].split('=')[1]
			}
        }
        return null;
    }



	$: cookieAccessToken = getCookie("access_token");
</script>

<svelte:head>
	<link
		rel="stylesheet"
		href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.7.0/font/bootstrap-icons.css"
	/>
	<link
		rel="stylesheet"
		href="https://cdn.jsdelivr.net/npm/bootstrap@5.1.3/dist/css/bootstrap.min.css"
	/>
</svelte:head>

{#if cookieAccessToken != null}
	<Navbar />
{:else}
	<Navbar /> <Login />
{/if}

<style></style>
