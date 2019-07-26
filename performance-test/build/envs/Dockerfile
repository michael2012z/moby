FROM busybox
ENV    abc=def def="hello world"
RUN    [ "$abc,$def" = "def,hello world" ]
ENV    def=hello\ world v1=abc v2="hi there" v3='boogie nights' v4="with'quotes too"
RUN    [ "$def,$v1,$v2,$v3,$v4" = "hello world,abc,hi there,boogie nights,with'quotes too" ]
ENV    abc 'yyy'
RUN    [ $abc = 'yyy' ]
ENV    abc=
RUN    [ "$abc" = "" ]

# use grep to make sure if the builder substitutes \$foo by mistake
# we don't get a false positive
ENV    abc=\$foo
RUN    [ "$abc" = "\$foo" ] && (echo "$abc" | grep foo)
ENV    abc \$foo
RUN    [ "$abc" = "\$foo" ] && (echo "$abc" | grep foo)

ENV    abc=\'foo\' abc2=\"foo\"
RUN    [ "$abc,$abc2" = "'foo',\"foo\"" ]
ENV    abc "foo"
RUN    [ "$abc" = "foo" ]
ENV    abc 'foo'
RUN    [ "$abc" = 'foo' ]
ENV    abc \'foo\'
RUN    [ "$abc" = "'foo'" ]
ENV    abc \"foo\"
RUN    [ "$abc" = '"foo"' ]

ENV    abc=ABC
RUN    [ "$abc" = "ABC" ]
ENV    def1=${abc:-DEF} def2=${ccc:-DEF}
ENV    def3=${ccc:-${def2}xx} def4=${abc:+ALT} def5=${def2:+${abc}:} def6=${ccc:-\$abc:} def7=${ccc:-\${abc}:}
RUN    [ "$def1,$def2,$def3,$def4,$def5,$def6,$def7" = 'ABC,DEF,DEFxx,ALT,ABC:,$abc:,${abc:}' ]
ENV    mypath=${mypath:+$mypath:}/home
ENV    mypath=${mypath:+$mypath:}/away
RUN    [ "$mypath" = '/home:/away' ]

ENV    e1=bar
ENV    e2=$e1 e3=$e11 e4=\$e1 e5=\$e11
RUN    [ "$e0,$e1,$e2,$e3,$e4,$e5" = ',bar,bar,,$e1,$e11' ]

ENV    ee1 bar
ENV    ee2 $ee1
ENV    ee3 $ee11
ENV    ee4 \$ee1
ENV    ee5 \$ee11
RUN    [ "$ee1,$ee2,$ee3,$ee4,$ee5" = 'bar,bar,,$ee1,$ee11' ]

ENV    eee1="foo" eee2='foo'
ENV    eee3 "foo"
ENV    eee4 'foo'
RUN    [ "$eee1,$eee2,$eee3,$eee4" = 'foo,foo,foo,foo' ]
