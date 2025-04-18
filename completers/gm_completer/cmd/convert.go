package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert an image or sequence of images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(convertCmd).Standalone()

	convertCmd.Flags().CountS("adjoin", "adjoin", "join images into a single multi-image file")
	convertCmd.Flags().StringSliceS("affine", "affine", []string{}, "affine transform matrix")
	convertCmd.Flags().CountS("antialias", "antialias", "remove pixel-aliasing")
	convertCmd.Flags().CountS("append", "append", "append an image sequence")
	convertCmd.Flags().StringSliceS("asc-cdl", "asc-cdl", []string{}, "apply ASC CDL transform")
	convertCmd.Flags().StringSliceS("authenticate", "authenticate", []string{}, "decrypt image with this password")
	convertCmd.Flags().CountS("auto-orient", "auto-orient", "orient (rotate) image so it is upright")
	convertCmd.Flags().CountS("average", "average", "average an image sequence")
	convertCmd.Flags().StringSliceS("background", "background", []string{}, "background color")
	convertCmd.Flags().StringSliceS("black-threshold", "black-threshold", []string{}, "pixels below the threshold become black")
	convertCmd.Flags().StringSliceS("blue-primary", "blue-primary", []string{}, "chomaticity blue primary point")
	convertCmd.Flags().StringSliceS("blur", "blur", []string{}, "blur the image")
	convertCmd.Flags().StringSliceS("border", "border", []string{}, "surround image with a border of color")
	convertCmd.Flags().StringSliceS("bordercolor", "bordercolor", []string{}, "border color")
	convertCmd.Flags().StringSliceS("box", "box", []string{}, "set the color of the annotation bounding box")
	convertCmd.Flags().StringSliceS("channel", "channel", []string{}, "extract a particular color channel from image")
	convertCmd.Flags().StringSliceS("charcoal", "charcoal", []string{}, "simulate a charcoal drawing")
	convertCmd.Flags().StringSliceS("chop", "chop", []string{}, "remove pixels from the image interior")
	convertCmd.Flags().CountS("clip", "clip", "apply first clipping path if the image has one")
	convertCmd.Flags().CountS("clippath", "clippath", "apply named clipping path if the image has one")
	convertCmd.Flags().CountS("coalesce", "coalesce", "merge a sequence of images")
	convertCmd.Flags().StringSliceS("colorize", "colorize", []string{}, "colorize the image with the fill color")
	convertCmd.Flags().CountS("colors", "colors", "preferred number of colors in the image")
	convertCmd.Flags().StringSliceS("colorspace", "colorspace", []string{}, "alternate image colorspace")
	convertCmd.Flags().StringSliceS("comment", "comment", []string{}, "annotate image with comment")
	convertCmd.Flags().StringSliceS("compose", "compose", []string{}, "composite operator")
	convertCmd.Flags().StringSliceS("compress", "compress", []string{}, "image compression type")
	convertCmd.Flags().CountS("contrast", "contrast", "enhance or reduce the image contrast")
	convertCmd.Flags().StringSliceS("convolve", "convolve", []string{}, "convolve image with the specified convolution kernel")
	convertCmd.Flags().StringSliceS("crop", "crop", []string{}, "preferred size and location of the cropped image")
	convertCmd.Flags().StringSliceS("cycle", "cycle", []string{}, "cycle the image colormap")
	convertCmd.Flags().StringSliceS("debug", "debug", []string{}, "display copious debugging information")
	convertCmd.Flags().CountS("deconstruct", "deconstruct", "break down an image sequence into constituent parts")
	convertCmd.Flags().StringSliceS("define", "define", []string{}, "Coder/decoder specific options")
	convertCmd.Flags().StringSliceS("delay", "delay", []string{}, "display the next image after pausing")
	convertCmd.Flags().StringSliceS("density", "density", []string{}, "horizontal and vertical density of the image")
	convertCmd.Flags().StringSliceS("depth", "depth", []string{}, "image depth")
	convertCmd.Flags().CountS("despeckle", "despeckle", "reduce the speckles within an image")
	convertCmd.Flags().StringSliceS("display", "display", []string{}, "get image or font from this X server")
	convertCmd.Flags().StringSliceS("dispose", "dispose", []string{}, "Undefined, None, Background, Previous")
	convertCmd.Flags().CountS("dither", "dither", "apply Floyd/Steinberg error diffusion to image")
	convertCmd.Flags().StringSliceS("draw", "draw", []string{}, "annotate the image with a graphic primitive")
	convertCmd.Flags().StringSliceS("edge", "edge", []string{}, "apply a filter to detect edges in the image")
	convertCmd.Flags().StringSliceS("emboss", "emboss", []string{}, "emboss an image")
	convertCmd.Flags().StringSliceS("encoding", "encoding", []string{}, "text encoding type")
	convertCmd.Flags().StringSliceS("endian", "endian", []string{}, "multibyte word order (LSB, MSB, or Native)")
	convertCmd.Flags().CountS("enhance", "enhance", "apply a digital filter to enhance a noisy image")
	convertCmd.Flags().CountS("equalize", "equalize", "perform histogram equalization to an image")
	convertCmd.Flags().CountS("extent", "extent", "composite image on background color canvas image")
	convertCmd.Flags().StringSliceS("fill", "fill", []string{}, "color to use when filling a graphic primitive")
	convertCmd.Flags().StringSliceS("filter", "filter", []string{}, "use this filter when resizing an image")
	convertCmd.Flags().CountS("flatten", "flatten", "flatten a sequence of images")
	convertCmd.Flags().CountS("flip", "flip", "flip image in the vertical direction")
	convertCmd.Flags().CountS("flop", "flop", "flop image in the horizontal direction")
	convertCmd.Flags().StringSliceS("font", "font", []string{}, "render text with this font")
	convertCmd.Flags().StringSliceS("format", "format", []string{}, "output formatted image info for 'info:' format")
	convertCmd.Flags().StringSliceS("frame", "frame", []string{}, "surround image with an ornamental border")
	convertCmd.Flags().StringSliceS("fuzz", "fuzz", []string{}, "colors within this distance are considered equal")
	convertCmd.Flags().StringSliceS("gamma", "gamma", []string{}, "level of gamma correction")
	convertCmd.Flags().StringSliceS("gaussian", "gaussian", []string{}, "gaussian blur an image")
	convertCmd.Flags().StringSliceS("geometry", "geometry", []string{}, "preferred size or location of the image")
	convertCmd.Flags().StringSliceS("gravity", "gravity", []string{}, "horizontal and vertical text/object placement")
	convertCmd.Flags().StringSliceS("green-primary", "green-primary", []string{}, "chomaticity green primary point")
	convertCmd.Flags().StringSliceS("hald-clut", "hald-clut", []string{}, "apply a Hald CLUT to the image")
	convertCmd.Flags().CountS("help", "help", "print program options")
	convertCmd.Flags().StringSliceS("implode", "implode", []string{}, "implode image pixels about the center")
	convertCmd.Flags().StringSliceS("intent", "intent", []string{}, "Absolute, Perceptual, Relative, or Saturation")
	convertCmd.Flags().StringSliceS("interlace", "interlace", []string{}, "None, Line, Plane, or Partition")
	convertCmd.Flags().StringSliceS("label", "label", []string{}, "assign a label to an image")
	convertCmd.Flags().StringSliceS("lat", "lat", []string{}, "local adaptive thresholding")
	convertCmd.Flags().StringSliceS("level", "level", []string{}, "adjust the level of image contrast")
	convertCmd.Flags().StringSliceS("limit", "limit", []string{}, "Disk, File, Map, Memory, Pixels, Width, Height, Threads, Read, or Write resource limit")
	convertCmd.Flags().StringSliceS("linewidth", "linewidth", []string{}, "the line width for subsequent draw operations")
	convertCmd.Flags().StringSliceS("list", "list", []string{}, "Color, Delegate, Format, Magic, Module, Resource, or Type")
	convertCmd.Flags().StringSliceS("log", "log", []string{}, "format of debugging information")
	convertCmd.Flags().StringSliceS("loop", "loop", []string{}, "add Netscape loop extension to your GIF animation")
	convertCmd.Flags().CountS("magnify", "magnify", "interpolate image to double size")
	convertCmd.Flags().StringSliceS("map", "map", []string{}, "transform image colors to match this set of colors")
	convertCmd.Flags().StringSliceS("mask", "mask", []string{}, "set the image clip mask")
	convertCmd.Flags().CountS("matte", "matte", "store matte channel if the image has one")
	convertCmd.Flags().StringSliceS("mattecolor", "mattecolor", []string{}, "specify the color to be used with the -frame option")
	convertCmd.Flags().StringSliceS("median", "median", []string{}, "apply a median filter to the image")
	convertCmd.Flags().CountS("minify", "minify", "interpolate the image to half size")
	convertCmd.Flags().StringSliceS("modulate", "modulate", []string{}, "vary the brightness, saturation, and hue")
	convertCmd.Flags().CountS("monitor", "monitor", "show progress indication")
	convertCmd.Flags().CountS("monochrome", "monochrome", "transform image to black and white")
	convertCmd.Flags().StringSliceS("morph", "morph", []string{}, "morph an image sequence")
	convertCmd.Flags().CountS("mosaic", "mosaic", "create a mosaic from an image sequence")
	convertCmd.Flags().StringSliceS("motion-blur", "motion-blur", []string{}, "simulate motion blur")
	convertCmd.Flags().CountS("negate", "negate", "replace every pixel with its complementary color ")
	convertCmd.Flags().StringSliceS("noise", "noise", []string{}, "add or reduce noise in an image")
	convertCmd.Flags().CountS("noop", "noop", "do not apply options to image")
	convertCmd.Flags().CountS("normalize", "normalize", "transform image to span the full range of colors")
	convertCmd.Flags().StringSliceS("opaque", "opaque", []string{}, "change this color to the fill color")
	convertCmd.Flags().StringSliceS("operator", "operator", []string{}, "apply a mathematical or bitwise operator to channel")
	convertCmd.Flags().StringSliceS("ordered-dither", "ordered-dither", []string{}, "ordered dither the image")
	convertCmd.Flags().StringSliceS("orient", "orient", []string{}, "set image orientation attribute")
	convertCmd.Flags().StringSliceS("page", "page", []string{}, "size and location of an image canvas")
	convertCmd.Flags().StringSliceS("paint", "paint", []string{}, "simulate an oil painting")
	convertCmd.Flags().CountS("ping", "ping", "efficiently determine image attributes")
	convertCmd.Flags().StringSliceS("pointsize", "pointsize", []string{}, "font point size")
	convertCmd.Flags().StringSliceS("preview", "preview", []string{}, "image preview type")
	convertCmd.Flags().StringSliceS("profile", "profile", []string{}, "add ICM or IPTC information profile to image")
	convertCmd.Flags().StringSliceS("quality", "quality", []string{}, "JPEG/MIFF/PNG compression level")
	convertCmd.Flags().StringSliceS("raise", "raise", []string{}, "lighten/darken image edges to create a 3-D effect")
	convertCmd.Flags().StringSliceS("random-threshold", "random-threshold", []string{}, "random threshold the image")
	convertCmd.Flags().StringSliceS("recolor", "recolor", []string{}, "apply a color translation matrix to image channels")
	convertCmd.Flags().StringSliceS("red-primary", "red-primary", []string{}, "chomaticity red primary point")
	convertCmd.Flags().StringSliceS("region", "region", []string{}, "apply options to a portion of the image")
	convertCmd.Flags().CountS("render", "render", "render vector graphics")
	convertCmd.Flags().StringSliceS("repage", "repage", []string{}, "adjust current page offsets by geometry")
	convertCmd.Flags().StringSliceS("resample", "resample", []string{}, "resample to horizontal and vertical resolution")
	convertCmd.Flags().StringSliceS("resize", "resize", []string{}, "resize the image")
	convertCmd.Flags().StringSliceS("roll", "roll", []string{}, "roll an image vertically or horizontally")
	convertCmd.Flags().StringSliceS("rotate", "rotate", []string{}, "apply Paeth rotation to the image")
	convertCmd.Flags().StringSliceS("sample", "sample", []string{}, "scale image with pixel sampling")
	convertCmd.Flags().StringSliceS("sampling-factor", "sampling-factor", []string{}, "horizontal and vertical sampling factors")
	convertCmd.Flags().StringSliceS("scale", "scale", []string{}, "scale the image")
	convertCmd.Flags().StringSliceS("scene", "scene", []string{}, "image scene number")
	convertCmd.Flags().StringSliceS("seed", "seed", []string{}, "pseudo-random number generator seed value")
	convertCmd.Flags().StringSliceS("segment", "segment", []string{}, "segment an image")
	convertCmd.Flags().StringSliceS("set", "set", []string{}, "set image attribute")
	convertCmd.Flags().StringSliceS("shade", "shade", []string{}, "shade the image using a distant light source")
	convertCmd.Flags().StringSliceS("sharpen", "sharpen", []string{}, "sharpen the image")
	convertCmd.Flags().StringSliceS("shave", "shave", []string{}, "shave pixels from the image edges")
	convertCmd.Flags().StringSliceS("shear", "shear", []string{}, "slide one edge of the image along the X or Y axis")
	convertCmd.Flags().StringSliceS("size", "size", []string{}, "width and height of image")
	convertCmd.Flags().StringSliceS("solarize", "solarize", []string{}, "negate all pixels above the threshold level")
	convertCmd.Flags().StringSliceS("spread", "spread", []string{}, "displace image pixels by a random amount")
	convertCmd.Flags().CountS("strip", "strip", "strip all profiles and text attributes from image")
	convertCmd.Flags().StringSliceS("stroke", "stroke", []string{}, "graphic primitive stroke color")
	convertCmd.Flags().StringSliceS("strokewidth", "strokewidth", []string{}, "graphic primitive stroke width")
	convertCmd.Flags().StringSliceS("swirl", "swirl", []string{}, "swirl image pixels about the center")
	convertCmd.Flags().StringSliceS("texture", "texture", []string{}, "name of texture to tile onto the image background")
	convertCmd.Flags().StringSliceS("threshold", "threshold", []string{}, "threshold the image")
	convertCmd.Flags().StringSliceS("thumbnail", "thumbnail", []string{}, "resize the image (optimized for thumbnails)")
	convertCmd.Flags().StringSliceS("tile", "tile", []string{}, "tile image when filling a graphic primitive")
	convertCmd.Flags().CountS("transform", "transform", "affine transform image")
	convertCmd.Flags().StringSliceS("transparent", "transparent", []string{}, "make this color transparent within the image")
	convertCmd.Flags().StringSliceS("treedepth", "treedepth", []string{}, "color tree depth")
	convertCmd.Flags().CountS("trim", "trim", "trim image edges")
	convertCmd.Flags().StringSliceS("type", "type", []string{}, "image type")
	convertCmd.Flags().StringSliceS("undercolor", "undercolor", []string{}, "annotation bounding box color")
	convertCmd.Flags().StringSliceS("units", "units", []string{}, "PixelsPerInch, PixelsPerCentimeter, or Undefined")
	convertCmd.Flags().StringSliceS("unsharp", "unsharp", []string{}, "sharpen the image")
	convertCmd.Flags().BoolS("verbose", "verbose", false, "print detailed information about the image")
	convertCmd.Flags().BoolS("version", "version", false, "print version information")
	convertCmd.Flags().CountS("view", "view", "FlashPix viewing transforms")
	convertCmd.Flags().StringSliceS("virtual-pixel", "virtual-pixel", []string{}, "Constant, Edge, Mirror, or Tile")
	convertCmd.Flags().StringSliceS("wave", "wave", []string{}, "alter an image along a sine wave")
	convertCmd.Flags().StringSliceS("white-point", "white-point", []string{}, "chomaticity white point")
	convertCmd.Flags().StringSliceS("white-threshold", "white-threshold", []string{}, "pixels above the threshold become white")
	convertCmd.Flags().StringSliceS("write", "write", []string{}, "write image to this file")
	rootCmd.AddCommand(convertCmd)

	carapace.Gen(convertCmd).FlagCompletion(carapace.ActionMap{
		"fill": action.ActionColor(),
	})

	carapace.Gen(convertCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
