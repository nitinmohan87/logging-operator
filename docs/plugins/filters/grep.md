# grep
## Overview
 #### Example grep filter configurations
 ```
 spec:
  filters:
    - regexp:
        - key: elso
          pattern: /^5\d\d$/
        - key: masodik
          pattern: /\.css$/
    - and:
        - regexp:
          - key: elso
            pattern: /^5\d\d$/
          exclude:
          - key: masodik
            pattern: /\.css$/
 ```

 #### Fluentd Config Result
 ```
<filter **>
  @type grep
  @id test_grep
  <regexp>
    key elso
    pattern /^5\d\d$/
  </regexp>
  <regexp>
    key masodik
    pattern /\.css$/
  </regexp>
  <and>
    <regexp>
      key elso
      pattern /^5\d\d$/
    </regexp>
    <exclude>
      key masodik
      pattern /\.css$/
    </exclude>
  </and>
</filter>
 ```

## Configuration
### GrepConfig
| Variable Name | Type | Required | Default | Description |
|---|---|---|---|---|
| regexp | []RegexpSection | No | - | [Regexp Section](#Regex-Section)<br> |
| exclude | []ExcludeSection | No | - | [Exclude Section](#Exclude-Section)<br> |
| or | []OrSection | No | - | [Or Section](#Or-Section)<br> |
| and | []AndSection | No | - | [And Section](#And-Section)<br> |
### Regexp Section
| Variable Name | Type | Required | Default | Description |
|---|---|---|---|---|
| key | string | Yes | - | Specify field name in the record to parse.<br> |
| pattern | string | Yes | - | Pattern expression to evaluate<br> |
### Exclude Section
| Variable Name | Type | Required | Default | Description |
|---|---|---|---|---|
| key | string | Yes | - | Specify field name in the record to parse.<br> |
| pattern | string | Yes | - | Pattern expression to evaluate<br> |
### Or directive
| Variable Name | Type | Required | Default | Description |
|---|---|---|---|---|
| regexp | []RegexpSection | No | - | [Regexp Section](#Regex-Section)<br> |
| exclude | []ExcludeSection | No | - | [Exclude Section](#Exclude-Section)<br> |
### And directive
| Variable Name | Type | Required | Default | Description |
|---|---|---|---|---|
| regexp | []RegexpSection | No | - | [Regexp Section](#Regex-Section)<br> |
| exclude | []ExcludeSection | No | - | [Exclude Section](#Exclude-Section)<br> |
