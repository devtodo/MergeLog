## MergeLog
---
### config.yml

```yaml
settings:
    monitoringTarget:
    targetFolder:
    tryCollectLog: true
    mergeLog: true
    convertPantaLabLogFormat: true
    mergedFileFormat: "★★★_merged_{folderName}.log"
    tryUpdate: true
    analysisEvent: true
    analysisLogPrefix: "=====@@@@@{{"
    analysisLogPostfix: "}}@@@@@====="
    analysisUpdater: false
    analysisProcessLifeCycle: true
    lastDayOnly: false
    deleteSourceFilesAfterCollect: true
    deleteSourceFilesAfterMerge: false
patterns:
    explorer:
        - /DoDragDrop Begin/: "드래그 & 드랍 시작"
        - /End DragNDrop/: "드래그 & 드랍 종료"
```
### Command Line Interface

```cli
MergeLog help
```