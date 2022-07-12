
import re
import os

log_pattern = re.compile(r"(.*)T([^\s]*)\s*\[([^]]*)\](.*)\ - (.*)")

with open("artifactory-service.log", "r") as f:
  for line in f:
      match = log_pattern.match(line)
      if not match:
        continue
      grps = match.groups()
      #print(f"  date:{grps[0]},\n  time:{grps[1]},\n  type:{grps[2]},\n  text:{grps[3]},\n data:{grps[4]}")
      #print(f"{grps[4]}")
      updaterepo = re.search(r'Updating repository (.*)', str(grps[4]), re.IGNORECASE)
      if updaterepo is None:
        pass
      else:
        print(f" Updated ==> " + str(updaterepo.group(1)))
    
      createrepo = re.search(r'Creating repository (.*)', str(grps[4]), re.IGNORECASE)
      if createrepo is None:
        pass
      else:
        print(f" Created ==> " + str(createrepo.group(1)))