package stringHTML

//WidgetTwitch ... strigifies HTML for widget
func WidgetTwitch(twitchUsername string) string {
	output := `
	<li class="ui-state-default widgetListItem sortable">
	<div class="projectsWidgetCont">
		<div class="widgetTitle">
		<span class="pull-right fa fa-2x fa-sort"></span>
		<span class="pull-right fa fa-2x fa-trash"></span>
		<span class="pull-right fa fa-2x fa-pencil" id="twitch-edit" data-toggle="modal" data-target="#twitch-modal"></span>
			<h4>Twitch</h4>
		</div>

		<div class="modal fade" id="twitch-modal" role="dialog">
			<div class="modal-dialog">
				<div class="modal-content">
					<div class="modal-header">
						<button type="button" class="close" data-dismiss="modal">&times;</button>
						<h4 class="modal-title">Embed Twitch</h4>
					</div>
					<div class="modal-body">
						<p>Put your Twitch username here to broadcast yourself whenever you stream live:</p>
						<div class="input-group">
							<div class="input-group-btn">
								<button id="twitch-setting" type="button" class="btn btn-secondary" aria-expanded="false">
									Username
								</button>
							</div>
							<input type="text" class="form-control" id="twitch-embed-username"/>
						</div>
					</div>
					<div class="modal-footer">
						<button type="submit" id="twitch-submit-btn" class="btn btn-default" data-dismiss="modal">Submit</button>
						<button type="button" class="btn btn-primary" data-dismiss="modal">Close</button>
					</div>
				</div>
			</div>
		</div> 

		<div id="widgetBodyTwitch" class="widgetBody">
			<iframe src="https://player.twitch.tv/?channel=' + ` + twitchUsername + ` + '" frameborder="0" allowfullscreen="true" scrolling="no" height="378" width="620"></iframe><a href="https://www.twitch.tv/' + ` + twitchUsername + ` + '?tt_content=text_link&tt_medium=live_embed" style="padding:2px 0px 4px; display:inline-block; width:345px; font-weight:normal; font-size:10px; text-decoration:underline;">Watch live video from ' ` + twitchUsername + ` ' on www.twitch.tv</a>
		</div>
	</div>
</li>
`
	return output
}
