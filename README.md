# see

A fast, lightweight command-line tool to preview markdown files as HTML in your browser or save them as HTML files with customizable templates.

## Why 'see'?

Writing Markdown for a README or documentation always turns into the same little struggle. You can’t see what anything looks like without setting up some preview tool, the styling is almost always plain and inflexible, and you end up constantly flipping between your editor and a browser just to check simple changes. And using an online editor never feels great either—you’re basically handing your content over to some random service. That’s why **see** feels so refreshing. It’s just one simple command-line tool that instantly opens your Markdown with a clean HTML preview, lets you switch between themes or use your own, and stays completely local so nothing ever leaves your machine. It ends up saving a ton of time, gives you professional-looking output without any hassle, works for everything from READMEs to blog posts, and stays super lightweight as a single tiny binary.

## Quick Start

Get started in 30 seconds:

```bash
# Install
go install github.com/zeusdijkstra/see@latest

# Preview any markdown file
see -file README.md

# Use dark theme
see -file README.md -t dark

# Save as HTML without preview
see -file README.md -s
```

## Installation

### From Source

```bash
git clone https://github.com/zeusdijkstra/see.git
cd see
go build -o see .
```

### Install Directly

```bash
go install github.com/zeusdijkstra/see@latest
```

## Usage

```bash
see -file <markdown_file> [-t <template>] [-s]
```

### Command Line Options

| Option | Description |
|--------|-------------|
| `-file` | Path to the markdown file to process (required) |
| `-t` | Template name (`default`, `minimal`, `dark`) or path to custom template file (optional) |
| `-s` | Skip browser preview and save HTML file instead (optional) |

## Examples

### Basic Usage
```bash
# Preview with default template
see -file README.md

# Use dark theme
see -file documentation.md -t dark

# Use minimal theme 
see -file article.md -t minimal
```

### Save HTML Files
```bash
# Save as HTML in current working directory without opening the browser
see -file notes.md -s

# Save with custom template
see -file blog.md -t /path/to/blog_template.html -s
```

### Custom Templates
```bash
# Use your own template
see -file report.md -t my_template.html

# Template with absolute path
see -file presentation.md -t /home/user/templates/slides.html
```

### Creating Custom Templates

Create custom templates using Go HTML template syntax with these variables:

- `{{.Title}}` - The document title (string)
- `{{.Body}}` - Sanitized HTML content (template.HTML)
- `{{.FileName}}` - Source file path (string)

**Error Handling**: If you specify an invalid template path, `see` will return error message like: `parse custom template "invalid.html": open invalid.html: no such file or directory`

#### Example Custom Template

Save this as `my_template.html`:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{.Title}}</title>
    <style>
        body { 
            font-family: 'Georgia', serif; 
            line-height: 1.8; 
            max-width: 800px; 
            margin: 0 auto; 
            padding: 2em;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
        }
        .content {
            background: rgba(255,255,255,0.1);
            padding: 2em;
            border-radius: 10px;
            backdrop-filter: blur(10px);
        }
        h1 { border-bottom: 2px solid rgba(255,255,255,0.3); }
        pre { background: rgba(0,0,0,0.3); padding: 1em; border-radius: 5px; }
    </style>
</head>
<body>
    <div class="content">
        <h1>{{.Title}}</h1>
        <p><small>Source: {{.FileName}}</small></p>
        <hr>
        {{.Body}}
    </div>
</body>
</html>
```

Usage: `see -file example.md -t my_template.html`

## How It Works

1. **Parsing**: Uses [Blackfriday v2](https://github.com/russross/blackfriday) to convert markdown to HTML
2. **Sanitization**: Applies [bluemonday](https://github.com/microcosm-cc/bluemonday) UGC policy for security
3. **Templating**: Renders HTML using Go's template engine with embedded or custom templates
4. **Preview**: Creates temporary HTML file using `mdp*.html` pattern and opens in default browser
5. **Cleanup**: Automatically removes temporary files after preview

## Browser Compatibility

The tool automatically detects your operating system and uses the appropriate command:

| Platform | Command |
|----------|---------|
| **Linux** | `xdg-open` |
| **macOS** | `open` |
| **Windows** | `rundll32 url.dll,FileProtocolHandler` |

## Testing

Run the test suite:

```bash
go test -v
```

The tests cover:
- All built-in templates
- Custom template functionality
- File output generation
- HTML normalization and comparison

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
