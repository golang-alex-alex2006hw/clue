#Clue - CLI Glue
**Clue** is a Go package that enables you to store persistent variables between application runs.  In certain cases like CLI interfaces, you may want to ensure that certain variables like authentication tokens are saved for later use.  The CLI application should be able to quickly store the information for immediate retrieval upon future runs consistently across OS platforms.


## Overview

This package does this by storing variables with the ```EncodeGobFile``` function with a ```UseValue``` struct.  The struct is stored in a Go-Binary file based on the ```os.TempDir()``` method and a parameter for suffix.  The ```TempDir``` is unique per OS login session and works across OS's.

Retrieval of parameters can be done with the ```DecodeGobFile``` function and a ```GetValue``` struct.  


## Examples
### EncodeGobFile
The following example creates an encoded Go-Binary file with a suffix of ```goair_compute```.  Following this, a ```clue.UseValue``` struct is created with a new ```map[string]string``` that defines extra parameters.  Specifically here we are storing an authentication token.

    err = clue.EncodeGobFile("goair_compute", clue.UseValue{
      VarMap: map[string]string{
        "VCDToken":           client.VCDToken,
      },
    })
    if err != nil {
      return fmt.Errorf("Error encoding gob: %s", err)
    }

### DecodeGobFile
It is easy to decode the GobFile.  Using the same suffix of ```goair_compute``` we would apply the GobFile to ```getValue``` struct.

    getValue := clue.GetValue{}
    if err := clue.DecodeGobFile("goair_compute", &getValue); err != nil {
      return fmt.Errorf("Problem with client DecodeGobFile", err)
    }

This then allows you to use the ```VarMap``` parameter to check if a reference exists and that then set the ```VCDToken``` to the string result.

    if getValue.VarMap["VCDToken"] != nil {
      VCDToken := *getValue.VarMap["VCDToken"]
    }



Licensing
---------
Licensed under the Apache License, Version 2.0 (the “License”); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an “AS IS” BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

Support
-------
Please file bugs and issues at the Github issues page. For more general discussions you can contact the EMC Code team at <a href="https://groups.google.com/forum/#!forum/emccode-users">Google Groups</a> or tagged with **EMC** on <a href="https://stackoverflow.com">Stackoverflow.com</a>. The code and documentation are released with no warranties or SLAs and are intended to be supported through a community driven process.
