package stringHTML

//WidgetInstagram ... strigifies HTML for widget
func WidgetInstagram(url string) string {
	output := `
<li class="ui-state-default widgetListItem sortable">
	<div class="projectsWidgetCont">
		<div class="widgetTitle">
			<span class="pull-right fa fa-2x fa-sort"></span>
			<span class="pull-right fa fa-2x fa-trash"></span>
			<span class="pull-right fa fa-2x fa-pencil" id="ig-edit" data-toggle="modal" data-target="#ig-modal"></span>
			<h4>Instagram</h4>
		</div>
		<div class="modal fade" id="ig-modal" role="dialog">
			<div class="modal-dialog">
				<div class="modal-content">
					<div class="modal-header">
						<button type="button" class="close" data-dismiss="modal">&times;</button>
						<h4 class="modal-title">Embed Instagram Posts</h4>
					</div>
					<form>
						<div class="modal-body">
							<span id="ig-list-title">Removable Instagram posts</span>
							<ul id="ig-edit-list"></ul>
							<label>Please paste the Instagram post here</label>
							<input type="text" class="form-control" id="ig-embed-input" placeholder="https://www.instagram.com/p/BageGbgD05c/" required/>
							<label>Captions</label>
							<input type="checkbox" class="form-check" id="ig-embed-caption" checked/>
						</div>
						<div class="modal-footer">
							<button type="submit" id="ig-submit-btn" class="btn btn-default">Submit</button>
						</div>
					</form>
				</div>
			</div>
		</div> 
		<div id="widgetBodyID" class="widgetBody">
			<div class="insta-feed"><blockquote class="instagram-media" data-instgrm-captioned data-instgrm-version="7" style=" background:#FFF; border:0; border-radius:3px; box-shadow:0 0 1px 0 rgba(0,0,0,0.5),0 1px 10px 0 rgba(0,0,0,0.15); margin: 1px; max-width:658px; padding:0; width:99.375%; width:-webkit-calc(100% - 2px); width:calc(100% - 2px);"><div style="padding:8px;"> <div style=" background:#F8F8F8; line-height:0; margin-top:40px; padding:50.0% 0; text-align:center; width:100%;"> <div style=" background:url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACwAAAAsCAMAAAApWqozAAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAAMUExURczMzPf399fX1+bm5mzY9AMAAADiSURBVDjLvZXbEsMgCES5/P8/t9FuRVCRmU73JWlzosgSIIZURCjo/ad+EQJJB4Hv8BFt+IDpQoCx1wjOSBFhh2XssxEIYn3ulI/6MNReE07UIWJEv8UEOWDS88LY97kqyTliJKKtuYBbruAyVh5wOHiXmpi5we58Ek028czwyuQdLKPG1Bkb4NnM+VeAnfHqn1k4+GPT6uGQcvu2h2OVuIf/gWUFyy8OWEpdyZSa3aVCqpVoVvzZZ2VTnn2wU8qzVjDDetO90GSy9mVLqtgYSy231MxrY6I2gGqjrTY0L8fxCxfCBbhWrsYYAAAAAElFTkSuQmCC); display:block; height:44px; margin:0 auto -44px; position:relative; top:-22px; width:44px;"></div></div> <p style=" margin:8px 0 0 0; padding:0 4px;"> <a href="' + ` + url + ` + '" style=" color:#000; font-family:Arial,sans-serif; font-size:14px; font-style:normal; font-weight:normal; line-height:17px; text-decoration:none; word-wrap:break-word;" target="_blank"></a></p> <p style=" color:#c9c8cd; font-family:Arial,sans-serif; font-size:14px; line-height:17px; margin-bottom:0; margin-top:8px; overflow:hidden; padding:8px 0 7px; text-align:center; text-overflow:ellipsis; white-space:nowrap;">A post shared by DIY CRAFTS FOOD LIFE HACKS ?? (@diy.learning) on <time style=" font-family:Arial,sans-serif; font-size:14px; line-height:17px;" datetime="2017-10-20T04:28:48+00:00">Oct 19, 2017 at 9:28pm PDT</time></p></div></blockquote> <script async defer src="//platform.instagram.com/en_US/embeds.js"></script></div>
		</div>
	</div>
</li>
`
	return output
}
