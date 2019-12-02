# The grep filter plugin "greps" events by the values of specified fields.
## Overview
 More info at https://docs.fluentd.org/filter/grep
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
| regexp | []RegexpSection | No | - | [Regexp Section](#Regex-Directive)<br> |
| exclude | []ExcludeSection | No | - | [Exclude Section](#Exclude-Directive)<br> |
| or | []OrSection | No | - | [Or Section](#Or-Directive)<br> |
| and | []AndSection | No | - | [And Section](#And-Directive)<br> |
### Regexp Directive
#### Specify filtering rule. This directive contains two parameters.
More info at https://docs.fluentd.org/filter/grep#less-than-regexp-greater-than-directive
#### Example Regexp filter configurations
```
spec:
filters:
- regexp:
- key: elso
pattern: /^5\d\d$/
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
</filter>
---
```

| Variable Name | Type | Required | Default | Description |
|---|---|---|---|---|
| key | string | Yes | - | Specify field name in the record to parse.<br> |
| pattern | string | Yes | - | Pattern expression to evaluate<br> |
### Exclude Directive
#### Specify filtering rule to reject events. This directive contains two parameters.
More info at https://docs.fluentd.org/filter/grep#less-than-exclude-greater-than-directive
#### Example Exclude filter configurations
```
spec:
filters:
- exclude:
- key: elso
pattern: /^5\d\d$/
```

#### Fluentd Config Result
```
<filter **>
@type grep
@id test_grep
<exclude>
key elso
pattern /^5\d\d$/
</exclude>
</filter>
```

| Variable Name | Type | Required | Default | Description |
|---|---|---|---|---|
| key | string | Yes | - | Specify field name in the record to parse.<br> |
| pattern | string | Yes | - | Pattern expression to evaluate<br> |
### Or Directive
#### Specify filtering rule. This directive contains either <regexp> or <exclude> directive.
More info at https://docs.fluentd.org/filter/grep#less-than-or-greater-than-directive
#### Example Or filter configurations
```
spec:
- or:
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
<or>
<regexp>
key elso
pattern /^5\d\d$/
</regexp>
<exclude>
key masodik
pattern /\.css$/
</exclude>
</or>
</filter>
```

| Variable Name | Type | Required | Default | Description |
|---|---|---|---|---|
| regexp | []RegexpSection | No | - | [Regexp Section](#Regex-Directive)<br> |
| exclude | []ExcludeSection | No | - | [Exclude Section](#Exclude-Directive)<br> |
### And Directive
#### Specify filtering rule. This directive contains either <regexp> or <exclude> directive.
More info at https://docs.fluentd.org/filter/grep#less-than-and-greater-than-directive
#### Example and filter configurations
```
spec:
filters:
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

| Variable Name | Type | Required | Default | Description |
|---|---|---|---|---|
| regexp | []RegexpSection | No | - | [Regexp Section](#Regex-Directive)<br> |
| exclude | []ExcludeSection | No | - | [Exclude Section](#Exclude-Directive)<br> |
