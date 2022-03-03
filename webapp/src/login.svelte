<script lang="ts">
	import { Col, Container, Row, Badge, Icon, Alert } from "sveltestrap";
</script>



<main>
	<h1>Welcome to FIRE CD</h1>
	<p>Please make login with some of the Oauth providers:</p>
	<div class="container-location">
		<Container>
			<Alert color='dark'  >
				<Row>
					<Col  md={{ size: 6, offset: 4 }}>
						<Row>
							<Col md="4">
								<h4>
									<a href="http://localhost:9090/api/repo/login/github">
										<Badge>Login with <Icon
											name="github"/>
										</Badge>
									</a>
								</h4>
							</Col>
						</Row>
						<Row>
							<Col md="4">
								<h4>
									<Badge
										>Login with <Icon
											name="github"
										/></Badge
									>
								</h4>
							</Col>
						</Row>
					</Col>
				</Row>
			</Alert>
		</Container>
	</div>
</main>

<style>
	.container-location {
		margin: auto;
		width: 40%;
	}
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>
