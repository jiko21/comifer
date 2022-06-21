# Comifer
git helper for generating commit log.

## overview
comifer is cli tool supporting for generating commit log.

Despite other tools, only you have to do is
1. install comifer
2. run `comifer init` to configure
3. write your own config(optional)

## config file
With config file, you can customize commit-log generation.
```json
{
  "format": "$1 $2",
  "steps": [
    {
      "type": "select",
      "message": "変更内容",
      "options": [
        {
          "value": ":bug:",
          "description": "bug fix"
        },
        {
          "value": ":+1:",
          "description": "plus one"
        }
      ]
    },
    {
      "type": "text",
      "message": "commit message"
    }
  ]
}
```
### format
`format` section specifies commit log format.
$1, $2, ... correspond to result of `steps` input.

### steps
Each `steps` item represents an interaction to generate commit log.

There are two types for `steps` item; `select` and `message`.

Each `steps` item has `message` field for asking question via cli.
#### select
`select` represents an interaction to choose items in `options`.

`options` item consits of two fields;
`description` and `value`.
| field | desc|
|-----:|:-----|
| value | value of selected option. this field is shown in log|
| description | description of value. this is shown in terminal |

#### text
`text` represents an action to fill text message for commit via terminal.

## How to install
For mac users, comifer can be installed via homebrew.
```bash
brew tap jiko21/comifer
brew install comifer
```
