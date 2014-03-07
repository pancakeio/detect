// heuristics determines how projects are built in git-based Pancake.io projects.
package heuristics

type StaticSiteType struct {
	Name    string
	Canary  string
	Command string
}

var DefaultBuild = &StaticSiteType{
	Name:    "html",
	Command: "cp -vLR $PANCAKE_SOURCE/* $PANCAKE_ARTIFACT_DIR",
}

var StaticSites = map[string]*StaticSiteType{
	"jekyll": {
		Name:    "jekyll",
		Canary:  "_config.yml",
		Command: "touch Gemfile && bundle install && bundle exec jekyll build --source $PANCAKE_SOURCE --destination $PANCAKE_ARTIFACT_DIR",
	},
	"pelican": {
		Name:    "pelican",
		Canary:  "pelicanconf.py",
		Command: "pelican $PANCAKE_SOURCE --output $PANCAKE_ARTIFACT_DIR --verbose",
	},
	"wintersmith": {
		Name:    "wintersmith",
		Canary:  "config.json",
		Command: "npm install && wintersmith build -C $PANCAKE_SOURCE -o $PANCAKE_ARTIFACT_DIR",
	},
	"middleman": {
		Name:    "middleman",
		Canary:  "config.rb",
		Command: "touch Gemfile && bundle install && bundle exec middleman build && cp -vLR $PANCAKE_SOURCE/build/* $PANCAKE_ARTIFACT_DIR/",
	},
	"hyde": {
		Name:    "hyde",
		Canary:  "info.yaml",
		Command: "hyde gen -s $PANCAKE_SOURCE -d $PANCAKE_ARTIFACT_DIR",
	},
	"sphinx": {
		Name:    "sphinx",
		Canary:  "conf.py",
		Command: "sphinx-build -b html $PANCAKE_SOURCE $PANCAKE_ARTIFACT_DIR",
	},
}
