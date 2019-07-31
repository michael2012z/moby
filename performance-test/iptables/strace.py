def strace(fileName):
    fin = open(fileName, 'r')
    inLines = fin.read().split('\n')
    fin.close()
    outLines = []
    for inLine in inLines:
        outLine = ''
        if len(inLine) == 0:
            continue
        elif inLine[-1] == '+':
            outLine = "------------------------"
            print(outLine)
        elif inLine[-1] == '>':
            words = inLine.split()
            outLine = words[1].split("(")[0] + "\t" + words[0] + "\t" +  words[-1][1:-1]
            print(outLine)
        else:
            continue
        outLines.append(outLine)
    if len(outLines) % 91 != 0:
        print("Something is wrong !!!")
    fout = open("time." + fileName, 'w')
    fout.write("\n".join(outLines))
    fout.close()

strace("strace.vm.txt")
