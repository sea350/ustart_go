package stringHTML

//WidgetGallery ... strigifies HTML for a gallery widget
func WidgetGallery(pictures []string) string {
	var output string
	if len(pictures) < 1 {
		return output
	}

	output = `
	<li class="ui-state-default widgetListItem sortable">
	<div class="projectsWidgetCont">
		<div class="widgetTitle">
			<span class="pull-right fa fa-2x fa-sort"></span>
			<span class="pull-right fa fa-2x fa-trash"></span>
			<span class="pull-right fa fa-2x fa-pencil" data-toggle="modal" data-target="#editGalleryModal"></span>
			<span id="galleryPause" class="pull-right fa fa-2x fa-pause"></span>
			<span id="galleryPlay" class="pull-right fa fa-2x fa-play"></span>
			<h4>Gallery</h4>
		</div>
		<div class="widgetBody">
			<div class="gallery-display">
				<!-- place the highlight image here -->
				<div><img src="` + pictures[0] + `" /></div>
			</div>
			<div class="gallery-slider slider-nav">
			   <!-- start a for loop here for the rest of the images -->`
	if len(pictures) > 1 {
		for _, img := range pictures {
			output += `
				<div><img src="` + img + `" /></div>
				<!-- end for loop-->
				`
		}
	}
	output += `
			</div>
		</div>
	</div>
	<div class="modal fade" id="editGalleryModal" role="dialog">
		<div class="modal-dialog">
			<div class="modal-content">
				<div class="modal-header">
					<button type="button" class="close" data-dismiss="modal">&times;</button>
					<h3>Gallery</h3>
				</div>
				<div class="modal-body">
					<label for="galleryURL">Add Pictures</label>
					<div class="input-group">
						<input id="galleryURL" class="form-control" type="text" placeholder="Place your image URL here"/>
						<span class="input-group-btn">
							<button id="galleryAdd" class="btn btn-default"><i class="fa fa-plus"></i></button>
						</span>
					</div>
					<label for="gallerySelect">Removes Pictures</label>
					<div class="input-group">
						<select id="gallerySelect" class="form-control" type="text"></select>
						<span class="input-group-btn">
							<button id="gallerySub" class="btn btn-default"><i class="fa fa-minus"></i></button>
						</span>
					</div>
				</div>
			</div>
		</div>
	</div>
</li>
	`
	return output
}
