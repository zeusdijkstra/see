# see

A fast, lightweight command-line tool to preview markdown files as HTML in your browser or save them as HTML files with customizable templates.

## Features

-  Instant browser preview of markdown files
-  Three built-in templates: default, minimal, and dark themes
-  Custom template support with Go HTML template syntax
-  Responsive design for all templates
-  HTML sanitization for security
-  Save HTML files without preview
-  Comprehensive test coverage

## Installation

### From Source

Clone the repository and build:

```bash
git clone <repository-url>
cd see
go build -o see .
```

### Install Directly

```bash
go install .
```

## Usage

```bash
see -file <markdown_file> [-t <template>] [-s]
```

### Command Line Options

- `-file`: Path to the markdown file to process (required)
- `-t`: Template name (`default`, `minimal`, `dark`) or path to custom template file (optional)
- `-s`: Skip browser preview and save HTML file instead (optional)

## Templates

### Built-in Templates

#### Default Template
- Clean, modern design with Segoe UI font
- Light gray background with white content area
- Responsive layout with mobile support
- Syntax highlighting for code blocks

#### Minimal Template
- Simple, distraction-free design
- Helvetica Neue font family
- Clean typography with subtle borders
- Perfect for focused reading

#### Dark Template
- Dark theme with #1a1a1a background
- High contrast for comfortable reading
- Blue accent colors (#4fc3f7)
- Optimized for low-light environments

### Creating Custom Templates

To create a custom template, use Go HTML template syntax with these available variables:

- `{{.Title}}` - The document title (string)
- `{{.Body}}` - Sanitized HTML content (template.HTML)
- `{{.FileName}}` - Source file path (string)

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

## Examples

### Basic Preview
```bash
see -file README.md
```
Opens README.md in your browser using the default template.

### Use Dark Theme
```bash
see -file documentation.md -t dark
```
Opens documentation.md with the dark theme.

### Save HTML Without Preview
```bash
see -file notes.md -s
```
Saves notes.md as notes.html in the current directory.

### Use Custom Template
```bash
see -file blog.md -t /path/to/blog_template.html
```
Opens blog.md using your custom template.

### Minimal Theme for Reading
```bash
see -file article.md -t minimal
```
Opens article.md with the minimal, distraction-free template.

## How It Works

1. **Parsing**: Uses [Blackfriday v2](https://github.com/russross/blackfriday) to convert markdown to HTML
2. **Sanitization**: Applies [bluemonday](https://github.com/microcosm-cc/bluemonday) UGC policy for security
3. **Templating**: Renders HTML using Go's template engine with embedded or custom templates
4. **Preview**: Creates temporary HTML file and opens in default browser
5. **Cleanup**: Automatically removes temporary files after preview

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

## Browser Compatibility

The tool automatically detects your operating system and uses the appropriate command:
- **Linux**: `xdg-open`
- **macOS**: `open`
- **Windows**: `rundll32 url.dll,FileProtocolHandler`

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
