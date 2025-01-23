// *** WARNING: this file was generated by pulumi-language-java. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.kubernetescertmanager.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CertManagerGlobalRbacArgs extends com.pulumi.resources.ResourceArgs {

    public static final CertManagerGlobalRbacArgs Empty = new CertManagerGlobalRbacArgs();

    @Import(name="create")
    private @Nullable Output<Boolean> create;

    public Optional<Output<Boolean>> create() {
        return Optional.ofNullable(this.create);
    }

    private CertManagerGlobalRbacArgs() {}

    private CertManagerGlobalRbacArgs(CertManagerGlobalRbacArgs $) {
        this.create = $.create;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CertManagerGlobalRbacArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CertManagerGlobalRbacArgs $;

        public Builder() {
            $ = new CertManagerGlobalRbacArgs();
        }

        public Builder(CertManagerGlobalRbacArgs defaults) {
            $ = new CertManagerGlobalRbacArgs(Objects.requireNonNull(defaults));
        }

        public Builder create(@Nullable Output<Boolean> create) {
            $.create = create;
            return this;
        }

        public Builder create(Boolean create) {
            return create(Output.of(create));
        }

        public CertManagerGlobalRbacArgs build() {
            return $;
        }
    }

}
